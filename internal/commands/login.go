package commands

import (
	"fmt"
	"net"
)

var (
	users map[string]string
)

func Login(userName string, conn net.Conn) {
	fmt.Printf("LOGIN: '%s'\n", userName)
	users[userName] = "#{userName} logged in"
	conn.Write([]byte("login successful"))
}
