package main

import (
	"io"
	"log"
	"math"
	"net"
	"os"
	"rpc_demo_server/pkg/protocol"
	"strconv"
	"time"
)

// The default listening port.
const PORT = 8333

// The default timeout duration in seconds.
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

	// Listen to the specified prot.
	listener, _ := net.Listen("tcp", ":"+strconv.Itoa(PORT))
	defer listener.Close()
	InfoLogger.Printf("Start to listen on port %d...\n", PORT)

	for {
		connection, _ := listener.Accept()

		InfoLogger.Printf(
			"Connected to %s\n",
			connection.RemoteAddr().String())

		// Use `goroutine` to handle client connection concurrently.
		go handleRequest(connection)
	}
}

// When a TCP connection was established between server and client,
// `handleRequest()` will be responsible to serve it till its termination.
func handleRequest(conn net.Conn) {
	defer conn.Close()

	// Set up connection time-out.
	conn.SetDeadline(time.Now().Add(
		TIMEOUT_DURATION * time.Second))

	// Keeps looping to read and process new requests,
	// until it's terminated by the client or time-out reached.
	for {
		// Read the head of the message.
		head := make([]byte, protocol.MSG_HEADER_SIZE)
		_, err := conn.Read(head)
		if err != nil {
			handleErr(conn, err)
			break
		}

		// Read the message payload.
		payload_len := protocol.GetPayloadLength(head[0])
		payload := make([]byte, payload_len)
		_, err = conn.Read(payload)
		if err != nil {
			handleErr(conn, err)
			break
		}

		// Read the checksum byte of the message.
		checksum := make([]byte, protocol.MSG_CHECKSUM_SIZE)
		_, err = conn.Read(checksum)
		if err != nil {
			handleErr(conn, err)
			break
		}

		InfoLogger.Printf(
			"Message with head %X and checksum %X (in hex) received",
			head, checksum)
		DebugLogger.Printf("Payload: %X\n", payload)

		// Form the received bytes into a `protocol.Message`.
		msg, err := protocol.NewMessageFromBytes(
			head[0],
			payload,
			checksum[0])
		if err != nil {
			ErrorLogger.Println("The message is invalid")
			ErrorLogger.Println(err.Error())
			ErrorLogger.Panicln("TODO")
		}

		processMessage(msg, conn)
	}

	InfoLogger.Printf(
		"Disconnected with %s\n",
		conn.RemoteAddr().String())
}

// handle all kinds of errors during reading from the TCP stream.
func handleErr(conn net.Conn, err error) {
	// reset the timeout duration
	// otherwise writting to client is disabled
	conn.SetDeadline(time.Time{})
	if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
		WarnLogger.Println("Connection timed out, closing")
		conn.Write(provideErrorMessage("Connection timeout..."))
	} else if err == io.EOF {
		InfoLogger.Println("Connection closed by client")
	} else {
		ErrorLogger.Println("Error reading: ", err.Error())
		conn.Write(provideErrorMessage("Error at reading..."))
	}
	// Try to make sure the write-back message can reach.
	time.Sleep(100 * time.Millisecond)
}

// Process a `protocol.Message` according to its message type.
func processMessage(msg protocol.Message, conn net.Conn) {
	switch protocol.GetMessageCode(msg) {
	case protocol.REQUEST:
		serialized_msg, err := handleMessageRequest(msg)
		if err != nil {
			ErrorLogger.Fatal(err.Error())
		}
		conn.Write(serialized_msg)
		InfoLogger.Println("Response sent")

	case protocol.ERROR:
		handleMessageError(msg)

	default:
		handleMessageOther(msg)
	}
}

// Handle a message with type `REQUEST`.
// Basically parse the floats and return calculated value.
func handleMessageRequest(msg protocol.Message) ([]byte, error) {
	// Parse received `Message` payload.
	payload := protocol.GetMessagePayload(msg)
	sides, err := protocol.ParsePayloadToFloat64s(payload)
	if err != nil {
		log.Panic(err.Error())
	}
	// Calculate the hypotenuse.
	DebugLogger.Println("Parse the sides as", sides)
	a := sides[0]
	b := sides[1]
	c := calculateHypotenuse(a, b)
	DebugLogger.Println("Get hypotenuse", c)
	// Return a new message as response, it should be serialized.
	return protocol.SerializeMessage(protocol.NewResponseMessage(c)), nil
}

func calculateHypotenuse(a float64, b float64) float64 {
	return math.Hypot(a, b)
}

func handleMessageError(msg protocol.Message) {
}

func handleMessageOther(msg protocol.Message) {
}

func provideErrorMessage(str string) []byte {
	panic("TODO")
}
