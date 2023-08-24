package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	// listen on port 8000
	ln, _ := net.Listen("tcp", ":8000")

	// accept connection
	conn, _ := ln.Accept()

	message, _ := bufio.NewReader(conn).ReadString('\n')

	println("Hello Server!")

	fmt.Printf("% X\n", message)
	conn.Write([]byte(message))
}
