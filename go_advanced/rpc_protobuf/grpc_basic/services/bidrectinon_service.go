package services

import (
	"fmt"
	"log"

	"github.com/x14n/goExperimental/go_advanced/rpc_protobuf/grpc_basic/protos"
)

type Bidirection struct {
	protos.UnimplementedBidireactionalServer
}

func (bd *Bidirection) BidirectionalHello(server protos.Bidireactional_BidirectionalHelloServer) error {
	defer func() {
		fmt.Println("client break connection")
	}()

	for {
		recv, err := server.Recv()
		if err != nil {
			return err
		}
		log.Println(recv)
		//send server message
		err = server.Send(&protos.BiResponse{Message: "fuck u client"})
		if err != nil {
			return err
		}
	}
}
