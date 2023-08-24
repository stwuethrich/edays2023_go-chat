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
	echo(connection)
	login(connection)
	listUsers(connection)
	logout(connection)
}

func login(connection net.Conn) {
	protocol.SendString(0, connection, "LOGIN goly")
	printAnswer(connection)
}

func listUsers(connection net.Conn) {
	protocol.SendString(0, connection, "USERS")
	printAnswer(connection)
}
func logout(connection net.Conn) {
	protocol.SendString(0, connection, "LOGOUT goly")
	printAnswer(connection)
}

func echo(connection net.Conn) {
	_, err := protocol.SendString(0, connection, "ECHO hello!")
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
