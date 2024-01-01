package main

import (
	"fmt"
	"io"
	"net"
	"rpc_demo/server/pkg/protocol"
	"strconv"
	"time"
)

var PORT = 8333

func main() {
	fmt.Printf("packages worked %f\n", protocol.GetFloat())
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
	conn.SetDeadline(time.Now().Add(10 * time.Second))

	for {
		// Read the head of the message
		head := make([]byte, 1)
		_, err := conn.Read(head)
		if err != nil {
			handleErr(conn, err)
			break
		}

		// Calculate the message length accroding to the protocol definition
		msg_len := head[0] & 0x3F

		// Read the message payload
		payload_len := msg_len - 2
		payload := make([]byte, payload_len)
		_, err = conn.Read(payload)
		if err != nil {
			handleErr(conn, err)
			break
		}

		// Read the checksum bit of the message
		checksum := make([]byte, 1)
		_, err = conn.Read(checksum)
		if err != nil {
			handleErr(conn, err)
			break
		}

		conn.Write([]byte("Message received.\n"))

		fmt.Println("Received (in hex):")
		fmt.Printf("\thead:   \t%x\n", head)
		fmt.Printf("\tpayload:\t%x\n", payload)
		fmt.Printf("\tchecksum:\t%x\n", checksum)
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
