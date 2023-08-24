package commands

import (
	"fmt"
	"net"
)

func EchoCommand(command string, message string, conn net.Conn) {
	fmt.Printf("'%s': '%s'\n", command, message)
	conn.Write([]byte(message))
	fmt.Printf("Terminate server\n")
}
