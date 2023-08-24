package commands

import (
	"github.com/stwuethrich/edays2023_go-chat/internal/protocol"
	"strings"
)

var (
	groups = make(map[string]string)
)

func Group(id int, groupMessage string, s Sender) (err error) {
	protocol.Log(id, "GROUP: '%s'\n", groupMessage)
	splitGroupMessage := strings.Split(groupMessage, " ")
	cmd, name := splitGroupMessage[0], splitGroupMessage[1]
	switch cmd {
	case "create":
		return create(id, name, s)
	case "remove":
		return remove(id, name, s)
	}
	return
}

func create(id int, name string, s Sender) (err error) {
	if _, groupExists := groups[name]; groupExists {
		protocol.SendToChannel(id, s.Ch, "group already created")
	} else {
		groups[name] = "#{name} created"
		defer protocol.Log(id, "group creation done")
		protocol.SendToChannel(id, s.Ch, "group created")
	}
	return err
}

func remove(id int, name string, s Sender) (err error) {
	if _, groupExists := groups[name]; groupExists {
		delete(groups, name)
		defer protocol.Log(id, "group removed")
		protocol.SendToChannel(id, s.Ch, "group removed")
	} else {
		protocol.SendToChannel(id, s.Ch, "group doesn't exist")
	}
	return err
}
