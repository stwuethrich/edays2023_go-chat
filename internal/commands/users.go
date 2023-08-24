package commands

import (
	"errors"
	"net"
)

type sender struct {
	conn net.Conn
	ch   chan string
}

func (s sender) send() (err error) {
	for err == nil {
		message, ok := <-s.ch
		if ok {
			_, err = s.conn.Write([]byte(message))
		} else {
			err = errors.New("chanel closed")
		}
	}
	return err
}

var (
	users = make(map[string]sender)
)
