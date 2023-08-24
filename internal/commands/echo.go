package commands

import (
	"github.com/stwuethrich/edays2023_go-chat/internal/protocol"
)

func Echo(id int, message string, s Sender) (err error) {
	protocol.Log(id, "send '%s'", message)
	defer protocol.Log(id, "echo done")
	protocol.SendToChannel(id, s.Ch, message)
	return err
}
