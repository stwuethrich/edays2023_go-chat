package commands

import (
	"github.com/stwuethrich/edays2023_go-chat/internal/protocol"
	"net"
)

var (
	users = make(map[string]string)
)

func Login(id int, userName string, conn net.Conn) (err error) {
	protocol.Log(id, "LOGIN: '%s'\n", userName)
	if _, userExists := users[userName]; userExists {
		_, err = protocol.SendString(id, conn, "user already logged in")
	} else {
		users[userName] = "#{userName} logged in"
		defer protocol.Log(id, "login done")
		_, err = protocol.SendString(id, conn, "login successful")
	}
	return err
}
