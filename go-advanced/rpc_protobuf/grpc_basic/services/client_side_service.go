package services

import (
	"log"

	"github.com/x14n/goExperimental/go_advanced/rpc_protobuf/grpc_basic/protos"
)

type ClientSideService struct {
	protos.UnimplementedClientSideServer
}

func (c *ClientSideService) ClientSideHello(server protos.ClientSide_ClientSideHelloServer) error {
	for i := 1; i < 5; i++ {
		recv, err := server.Recv()
		if err != nil {
			return err
		}
		log.Println("client message:", recv)
	}
	err := server.SendAndClose(&protos.ClientResponse{Resp: "fuck u!"})
	if err != nil {
		return err
	}
	return nil
}
