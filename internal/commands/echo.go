package commands

import (
	"fmt"
	"github.com/stwuethrich/edays2023_go-chat/internal/protocol"
	"net"
)

func Echo(message string, conn net.Conn) {
	fmt.Printf("ECHO: '%s'\n", message)
	defer fmt.Println("echo done")
	protocol.SendString(conn, message)
}
