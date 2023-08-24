package commands

import (
	"github.com/stwuethrich/edays2023_go-chat/internal/protocol"
	"net"
)

func List(id int, message string, conn net.Conn) (err error) {
	protocol.Log(id, "USERS")
	var result string
	switch message {
	case "users":
		for key, _ := range users {
			result += key + ","
		}
	}
	_, err = protocol.SendString(id, conn, result[:len(result)-1])
	return err
}
