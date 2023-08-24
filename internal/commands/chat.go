package commands

import (
	"github.com/stwuethrich/edays2023_go-chat/internal/protocol"
	"net"
)

func Chat(id int, message string, _ net.Conn) (err error) {
	for user, dest := range users {
		protocol.Log(id, "  %s (%d) <- %s", user, dest, message)
	}
	return err
}
