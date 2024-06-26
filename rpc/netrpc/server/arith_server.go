package server

import (
	"log"
	"net"
	"net/rpc"
)

type Args struct {
	A, B int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func ArithServer() {
	arith := new(Arith)
	rpc.Register(arith)
	liesten, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP: ", err)
	}

	for {
		conn, err := liesten.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}

		go rpc.ServeConn(conn)
	}
}
