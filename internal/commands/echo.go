package commands

import (
	"github.com/stwuethrich/edays2023_go-chat/internal/protocol"
	"net"
)

func Echo(id int, message string, conn net.Conn) (err error) {
	protocol.Log(id, "ECHO: '%s'\n", message)
	defer protocol.Log(id, "echo done")
	_, err = protocol.SendString(0, conn, message)
	return err
}
