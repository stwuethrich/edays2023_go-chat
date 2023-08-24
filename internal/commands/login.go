package commands

import (
	"github.com/stwuethrich/edays2023_go-chat/internal/protocol"
)

func Login(id int, userName string, s Sender) (err error) {
	protocol.Log(id, "LOGIN: '%s'\n", userName)
	if _, userExists := users[userName]; userExists {
		protocol.SendToChannel(id, s.Ch, "user already logged in")
	} else {
		users[userName] = s
		defer protocol.Log(id, "login done")
		protocol.SendToChannel(id, s.Ch, "login successful")
	}
	return err
}
