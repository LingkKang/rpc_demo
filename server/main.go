package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
)

var PORT = 8333

func main() {
	listener, _ := net.Listen("tcp", ":"+strconv.Itoa(PORT))
	defer listener.Close()
	fmt.Printf("Start to listen on port %d...\n\n", PORT)

	for {
		connection, _ := listener.Accept()
		go handleRequest(connection)
	}
}

func handleRequest(conn net.Conn) {
	defer conn.Close()

	buffer, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Print("Received message: ", string(buffer))

	conn.Write([]byte("Message received.\n"))
}
