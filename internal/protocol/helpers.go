package protocol

import (
	"fmt"
	"net"
)

func SendString(connection net.Conn, message string) (int, error) {
	message += "\n"
	fmt.Printf("sendString %s", message)
	return connection.Write([]byte(message))
}
