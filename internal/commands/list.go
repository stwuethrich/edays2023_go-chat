package commands

import (
	"github.com/stwuethrich/edays2023_go-chat/internal/protocol"
)

func List(id int, message string, s Sender) (err error) {
	protocol.Log(id, "USERS")
	var result string
	switch message {
	case "users":
		for key, _ := range users {
			result += key + ","
		}
	}
	protocol.SendToChannel(id, s.Ch, result[:len(result)-1])
	return err
}
