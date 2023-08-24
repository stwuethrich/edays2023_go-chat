package protocol

import (
	"fmt"
	"strconv"
)

func Log(id int, message string, param ...interface{}) {
	fmt.Printf(strconv.Itoa(id)+": "+message+"\n", param...)
}
