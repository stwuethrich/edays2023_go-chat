package commands

import (
	"github.com/stwuethrich/edays2023_go-chat/internal/protocol"
	"net"
)

func ListUsers(id int, conn net.Conn) (err error) {
	protocol.Log(id, "USERS")
	var result string
	for key, _ := range users {
		result += key + ","
	}
	_, err = protocol.SendString(id, conn, result[:len(result)-1])
	return err
}
