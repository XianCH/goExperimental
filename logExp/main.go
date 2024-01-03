package main

import (
	"fmt"
	"net"

	xlog "github.com/x14n/goExperimental/logExp/v2"
)

func main() {
	l, err := xlog.NewLogger("./", "HELLO.log")
	if err != nil {
		fmt.Println(err)
		return
	}

	c, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		l.INFO(err)
	}
	defer c.Close()
	defer l.Close()
}
