package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Printf("Start client\n")
	connection, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		panic(err)
	}
	///send some data
	fmt.Printf("Send to server\n")
	_, err = connection.Write([]byte("ECHO Hello Server! Greetings.\n"))
	buffer := make([]byte, 1024)
	fmt.Printf("Read from server\n")
	mLen, err := connection.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	fmt.Println("Received: ", string(buffer[:mLen]))
	defer connection.Close()
}
