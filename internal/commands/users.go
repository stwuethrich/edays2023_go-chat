package commands

import (
	"errors"
	"net"
)

type Sender struct {
	Conn net.Conn
	Ch   chan string
}

func (s Sender) Send() (err error) {
	for err == nil {
		message, ok := <-s.Ch
		if ok {
			_, err = s.Conn.Write([]byte(message))
		} else {
			err = errors.New("chanel closed")
		}
	}
	return err
}

var (
	users = make(map[string]Sender)
)
