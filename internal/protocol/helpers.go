package protocol

import (
	"net"
)

func SendString(id int, connection net.Conn, message string) (int, error) {
	Log(id, "sendString %s", message)
	message += "\n"
	return connection.Write([]byte(message))
}
