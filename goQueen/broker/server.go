package broker

import (
	"log"
	"net"
)

func StartBrokder() {
	broker := NewBroker()

	listen, err := net.Listen("tcp", "127.0.0.1:12345")
	if err != nil {
		log.Printf("StartBrokder err : %s\n", err)
		return
	}
	log.Println("broker server start", listen.Addr(), ".....")
	for {
		conn, err := listen.Accept()
		log.Println("get connect form", conn.RemoteAddr())
		if err != nil {
			return
		}
		go broker.Process(conn)
	}
}
