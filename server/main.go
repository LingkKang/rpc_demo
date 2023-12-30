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

	reader := bufio.NewReader(conn)

	for {
		message, _ := reader.ReadString('\n')
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
