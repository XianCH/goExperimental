package server

import (
	"log"
	"math/rand"
	"net"
	"net/rpc"
	"time"
)

type GuessArgs struct {
	Guess int
}

type GuessGameServer struct {
	Number int
	Low    int
	Hight  int
}

// useing algorithm to guess the number
func (t *GuessGameServer) GuessNumber(args *GuessArgs, reply *int) error {
	if args.Guess < t.Number {
		t.Low = args.Guess + 1
		*reply = -1
	} else if args.Guess > t.Number {
		t.Hight = args.Guess - 1
		*reply = 1
	} else {
		*reply = 0
	}
	return nil
}

func GuessStart() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	gameServer := &GuessGameServer{Number: r.Intn(100), Low: 0, Hight: 100}
	err := rpc.Register(gameServer)
	if err != nil {
		log.Printf("Register error: %v", err)
		return
	}
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Printf("Listen error: %v", err)
		return
	}
	for {
		conn, _ := listener.Accept()
		go rpc.ServeConn(conn)
	}
}
