package commands

import (
	"github.com/stwuethrich/edays2023_go-chat/internal/protocol"
)

func Logout(id int, userName string, s Sender) (err error) {
	protocol.Log(id, "LOGOUT: '%s'\n", userName)
	if _, userExists := users[userName]; userExists {
		delete(users, userName)
		defer protocol.Log(id, "logout done")
		protocol.SendToChannel(id, s.Ch, "logout successful")
	} else {
		protocol.SendToChannel(id, s.Ch, "user not logged in")
	}
	return err
}
