package commands

import (
	"github.com/stwuethrich/edays2023_go-chat/internal/protocol"
)

func List(id int, message string, s Sender) (err error) {
	protocol.Log(id, "USERS")
	var result string
	switch message {
	case "users":
		result = getKeys(users)
	case "groups":
		result = getKeys(groups)
	}
	protocol.SendToChannel(id, s.Ch, result)
	return err
}

func getKeys[T interface{}](m map[string]T) (result string) {
	for key, _ := range m {
		result += key + ","
	}
	l := len(result)
	if l > 0 {
		result = result[:l-1]
	}
	return
}
