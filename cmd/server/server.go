package main

import (
	"bufio"
	"fmt"
	"github.com/stwuethrich/edays2023_go-chat/internal/commands"
	"github.com/stwuethrich/edays2023_go-chat/internal/protocol"
	"net"
)

func main() {
	fmt.Printf("Start server\n")
	defer fmt.Printf("Terminate server\n")
	// listen on port 8000
	ln, _ := net.Listen("tcp", ":8000")

	for id := 0; ; id++ {
		// accept connection
		conn, _ := ln.Accept()
		protocol.Log(id, "connection accepted")
		go func(id int, conn2 net.Conn) {
			defer conn2.Close()
			var err error
			for err == nil {
				reader := bufio.NewReader(conn2)
				var command string
				command, err = reader.ReadString(' ')
				if err != nil {
					return
				}
				var message string
				message, err = reader.ReadString('\n')
				if err != nil {
					return
				}

				protocol.Log(id, "%s - %s", command, message)
				command = command[:len(command)-1]
				message = message[:len(message)-1]

				protocol.Log(id, "received command %s\n", command)
				switch command {
				case "ECHO":
					err = commands.Echo(id, message, conn2)
				case "LOGIN":
					err = commands.Login(id, message, conn2)
				case "LOGOUT":
					err = commands.Logout(id, message, conn2)
				case "USERS":
					err = commands.ListUsers(0, conn2)
				default:
					protocol.Log(id, "unknown command #{command}")
				}
			}
		}(id, conn)

	}
}
