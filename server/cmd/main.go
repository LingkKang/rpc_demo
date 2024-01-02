package main

import (
	"fmt"
	"io"
	"net"
	"rpc_demo/server/pkg/protocol"
	"strconv"
	"time"
)

const PORT = 8333
const TIMEOUT_DURATION = 30

func main() {
	listener, _ := net.Listen("tcp", ":"+strconv.Itoa(PORT))
	defer listener.Close()
	fmt.Printf("Start to listen on port %d...\n\n", PORT)

	for {
		connection, _ := listener.Accept()

		fmt.Printf(
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

		fmt.Println("Received (in hex):")
		fmt.Printf("\thead:   \t%X\n", head)
		fmt.Printf("\tpayload:\t%X\n", payload)
		fmt.Printf("\tchecksum:\t%X\n", checksum)

		msg, err := protocol.NewMessageFromBytes(
			head[0],
			payload,
			checksum[0])
		if err != nil {
			fmt.Println("The message is invalid")
			fmt.Println(err.Error())
			panic("TODO")
		}
		processMessage(msg)
	}

	fmt.Printf(
		"Disconnected with %s\n",
		conn.RemoteAddr().String())
}

func handleErr(conn net.Conn, err error) {
	// reset the timeout duration
	// otherwise writting to client is disabled
	conn.SetDeadline(time.Time{})
	if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
		fmt.Println("Connection timed out, closing")
		conn.Write([]byte("Connection timeout...\n"))
	} else if err == io.EOF {
		fmt.Println("Connection closed by client")
	} else {
		fmt.Println("Error reading: ", err.Error())
		conn.Write([]byte("Error at reading, closing...\n"))
	}
	time.Sleep(100 * time.Millisecond)
}

func processMessage(msg protocol.Message) {
	switch protocol.GetMessageCode(msg) {
	case protocol.REQUEST:
		handleRequestPayload(protocol.GetMessagePayload(msg))
	default:
		panic("Unimplemented message type\n")
	}
}

func handleRequestPayload(payload []byte) {
	panic("TODO")
}
