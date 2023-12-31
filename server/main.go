package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"
)

var PORT = 8333
var EXIT = "\xff\xf4\xff\xfd\x06" // Telnet interrupt signal

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

	conn.Write([]byte(
		"Greeting from Golang!\nType `exit` to exit\n"))

	conn.SetDeadline(time.Now().Add(10 * time.Second))

	reader := bufio.NewReader(conn)

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			// reset the timeout duration
			// otherwise writting to client is disabled
			conn.SetDeadline(time.Time{})
			if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
				fmt.Println("Connection timed out, closing")
				conn.Write([]byte("Connection timeout...\n"))
			} else {
				fmt.Println("Error reading: ", err.Error())
				conn.Write([]byte("Error at reading, closing...\n"))
			}
			time.Sleep(100 * time.Millisecond)
			break
		}
		trimed_message := trimMessage(message)

		fmt.Printf("Received message: %s\n", trimed_message)

		conn.Write([]byte("Message received.\n"))

		if trimed_message == "exit" || trimed_message == EXIT {
			conn.Write([]byte("Goodbye!\n"))
			time.Sleep(100 * time.Millisecond)
			break
		}
	}

	fmt.Printf(
		"Disconnected with %s\n",
		conn.RemoteAddr().String())
}

func trimMessage(str string) string {
	return strings.TrimRightFunc(
		str,
		func(r rune) bool {
			return r == ' ' || r == '\r' || r == '\n'
		})
}
