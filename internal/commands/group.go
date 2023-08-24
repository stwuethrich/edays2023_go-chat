package commands

import (
	"github.com/stwuethrich/edays2023_go-chat/internal/protocol"
	"net"
	"strings"
)

var (
	groups = make(map[string]string)
)

func Group(id int, groupMessage string, conn net.Conn) (err error) {
	protocol.Log(id, "GROUP: '%s'\n", groupMessage)
	splitGroupMessage := strings.Split(groupMessage, " ")
	cmd, name := splitGroupMessage[0], splitGroupMessage[1]
	switch cmd {
	case "create":
		return create(id, name, conn)
	case "remove":
		return remove(id, name, conn)
	}
	return
}

func create(id int, name string, conn net.Conn) (err error) {
	if _, groupExists := groups[name]; groupExists {
		_, err = protocol.SendString(id, conn, "group already created")
	} else {
		groups[name] = "#{name} created"
		defer protocol.Log(id, "group creation done")
		_, err = protocol.SendString(id, conn, "group created")
	}
	return err
}

func remove(id int, name string, conn net.Conn) (err error) {
	if _, groupExists := groups[name]; groupExists {
		delete(groups, name)
		defer protocol.Log(id, "group removed")
		_, err = protocol.SendString(id, conn, "group removed")
	} else {
		_, err = protocol.SendString(id, conn, "group doesn't exist")
	}
	return err
}
