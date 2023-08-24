package main

import (
	"bufio"
	"fmt"
	"github.com/stwuethrich/edays2023_go-chat/internal/protocol"
	"net"
)

func main() {
	fmt.Printf("Start client\n")
	connection, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		panic(err)
	} else {
		defer connection.Close()
	}

	///send some data
	fmt.Printf("Send to server\n")
	sendEcho(connection)
	login(connection)
}

func login(connection net.Conn) {
	protocol.SendString(connection, "LOGIN goly")
	printAnswer(connection)
}

func sendEcho(connection net.Conn) {
	_, err := protocol.SendString(connection, "ECHO hello!")
	if err != nil {
		fmt.Println("Error writing:", err.Error())
	}
	printAnswer(connection)
}

func printAnswer(connection net.Conn) {
	fmt.Println("reading...")
	reader := bufio.NewReader(connection)
	message, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	fmt.Println("Received: ", message)
}
