package commands

import (
	"github.com/stwuethrich/edays2023_go-chat/internal/protocol"
	"net"
)

func Echo(id int, message string, conn net.Conn) (err error) {
	protocol.Log(id, "send '%s'", message)
	defer protocol.Log(id, "echo done")
	_, err = protocol.SendString(id, conn, message)
	return err
}
