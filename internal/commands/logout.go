package commands

import (
	"github.com/stwuethrich/edays2023_go-chat/internal/protocol"
	"net"
)

func Logout(id int, userName string, conn net.Conn) (err error) {
	protocol.Log(id, "LOGOUT: '%s'\n", userName)
	if _, userExists := users[userName]; userExists {
		delete(users, userName)
		defer protocol.Log(id, "logout done")
		_, err = protocol.SendString(id, conn, "logout successful")
	} else {
		_, err = protocol.SendString(id, conn, "user not logged in")
	}
	return err
}
