package main

import (
	"io"
	"log"
	"math"
	"net"
	"os"
	"rpc_demo/server/pkg/protocol"
	"strconv"
	"time"
)

const PORT = 8333
const TIMEOUT_DURATION = 10

var (
	DebugLogger *log.Logger
	InfoLogger  *log.Logger
	WarnLogger  *log.Logger
	ErrorLogger *log.Logger
)

func main() {

	DebugLogger = log.New(os.Stdout, "[DEBUG]\t", log.Ldate|log.Ltime|log.LUTC)
	InfoLogger = log.New(os.Stdout, "[INFO]\t", log.Ldate|log.Ltime|log.LUTC)
	WarnLogger = log.New(os.Stdout, "[WARN]\t", log.Ldate|log.Ltime|log.LUTC)
	ErrorLogger = log.New(os.Stdout, "[ERROR]\t", log.Ldate|log.Ltime|log.LUTC)

	listener, _ := net.Listen("tcp", ":"+strconv.Itoa(PORT))
	defer listener.Close()
	InfoLogger.Printf("Start to listen on port %d...\n", PORT)

	for {
		connection, _ := listener.Accept()

		InfoLogger.Printf(
			"Connected to %s\n",
			connection.RemoteAddr().String())

		go handleRequest(connection)
	}
}

func handleRequest(conn net.Conn) {
	defer conn.Close()

	conn.Write([]byte("Greeting from Golang!\n"))
	conn.SetDeadline(time.Now().Add(
		TIMEOUT_DURATION * time.Second))

	for {
		// Read the head of the message
		head := make([]byte, protocol.MSG_HEADER_SIZE)
		_, err := conn.Read(head)
		if err != nil {
			handleErr(conn, err)
			break
		}

		// Read the message payload
		payload_len := protocol.GetPayloadLength(head[0])
		payload := make([]byte, payload_len)
		_, err = conn.Read(payload)
		if err != nil {
			handleErr(conn, err)
			break
		}

		// Read the checksum bit of the message
		checksum := make([]byte, protocol.MSG_CHECKSUM_SIZE)
		_, err = conn.Read(checksum)
		if err != nil {
			handleErr(conn, err)
			break
		}

		conn.Write([]byte("Message received.\n"))

		InfoLogger.Printf(
			"Message with head %X and checksum %X (in hex) received",
			head, checksum)
		DebugLogger.Printf("Payload: %X\n", payload)

		msg, err := protocol.NewMessageFromBytes(
			head[0],
			payload,
			checksum[0])
		if err != nil {
			ErrorLogger.Println("The message is invalid")
			ErrorLogger.Println(err.Error())
			ErrorLogger.Panicln("TODO")
		}
		processMessage(msg)
	}

	InfoLogger.Printf(
		"Disconnected with %s\n",
		conn.RemoteAddr().String())
}

func handleErr(conn net.Conn, err error) {
	// reset the timeout duration
	// otherwise writting to client is disabled
	conn.SetDeadline(time.Time{})
	if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
		WarnLogger.Println("Connection timed out, closing")
		conn.Write([]byte("Connection timeout...\n"))
	} else if err == io.EOF {
		InfoLogger.Println("Connection closed by client")
	} else {
		ErrorLogger.Println("Error reading: ", err.Error())
		conn.Write([]byte("Error at reading, closing...\n"))
	}
	time.Sleep(100 * time.Millisecond)
}

func processMessage(msg protocol.Message) {
	switch protocol.GetMessageCode(msg) {
	case protocol.REQUEST:
		handleMessageRequest(msg)
	case protocol.ERROR:
		handleMessageError(msg)
	default:
		handleMessageOther(msg)
	}
}

func handleMessageRequest(msg protocol.Message) {
	payload := protocol.GetMessagePayload(msg)
	sides, err := protocol.ParsePayloadToFloat(payload)
	if err == nil {
		DebugLogger.Println("Parse the sides as ", sides)
		a := sides[0]
		b := sides[1]
		c := calculateHypotenuse(a, b)
		DebugLogger.Println("Get hypotenuse ", c)
	}

}

func handleMessageError(msg protocol.Message) {

}

func handleMessageOther(msg protocol.Message) {

}

func calculateHypotenuse(a float64, b float64) float64 {
	return math.Hypot(a, b)
}
