package commands

import (
	"github.com/stwuethrich/edays2023_go-chat/internal/protocol"
)

func Chat(id int, message string) {
	for user, dest := range users {
		protocol.Log(id, "  %s <- %s", user, message)
		protocol.SendToChannel(id, dest.Ch, "CHAT "+user+": "+message)
	}
}
