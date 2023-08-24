package commands

import (
	"fmt"
	"net"
)

func Echo(message string, conn net.Conn) {
	fmt.Printf("ECHO: '%s'\n", message)
	conn.Write([]byte(message))
}
