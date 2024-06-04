package services

import (
	"log"

	"github.com/x14n/goExperimental/go_advanced/rpc_protobuf/grpc_basic/protos"
)

type ServerSideService struct {
	protos.UnimplementedServerSideServer
}

func (s *ServerSideService) ServerSideHello(req *protos.ServerSideRequest, server protos.ServerSide_ServerSideHelloServer) error {
	log.Println(req.GetSideReq())
	for n := 0; n < 5; n++ {
		err := server.Send(&protos.ServerSideResponse{SideRes: "hello!"})
		if err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}
