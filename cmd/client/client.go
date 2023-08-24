package main

import (
	"bufio"
	"fmt"
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
	sendString(connection, "LOGIN goly")
	printAnswer(connection)
}

func sendEcho(connection net.Conn) {
	_, err := sendString(connection, "ECHO hello!")
	if err != nil {
		fmt.Println("Error writing:", err.Error())
	}
	printAnswer(connection)
}

func printAnswer(connection net.Conn) {
	reader := bufio.NewReader(connection)
	message, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	fmt.Println("Received: ", message)
}

func sendString(connection net.Conn, message string) (int, error) {
	message += "\n"
	fmt.Printf("sendString %s", message)
	return connection.Write([]byte(message))
}
