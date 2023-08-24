package commands

import (
	"fmt"
	"github.com/stwuethrich/edays2023_go-chat/internal/protocol"
	"net"
)

var (
	users map[string]string
)

func Login(userName string, conn net.Conn) {
	fmt.Printf("LOGIN: '%s'\n", userName)
	users[userName] = "#{userName} logged in"
	defer fmt.Println("login done")
	protocol.SendString(conn, "login successful")
}
