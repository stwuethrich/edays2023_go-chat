package main

import (
	"bufio"
	"fmt"
	"github.com/stwuethrich/edays2023_go-chat/internal/commands"
	"net"
)

func main() {
	fmt.Printf("Start server\n")
	defer fmt.Printf("Terminate server\n")
	// listen on port 8000
	ln, _ := net.Listen("tcp", ":8000")

	for {
		// accept connection
		conn, _ := ln.Accept()

		reader := bufio.NewReader(conn)
		command, _ := reader.ReadString(' ')
		message, _ := reader.ReadString('\n')

		command = command[:len(command)-1]
		message = message[:len(message)-1]

		fmt.Println("received command #{command}")
		switch command {
		case "ECHO":
			go commands.Echo(message, conn)
		case "LOGIN":
			go commands.Login(message, conn)
		default:
			fmt.Println("unknown command #{command}")
		}

	}
}
