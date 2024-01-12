package main

import (
	"log"
	"net"

	"github.com/x14n/goExperimental/go_advanced/rpc_protobuf/grpc_basic/protos"
	"github.com/x14n/goExperimental/go_advanced/rpc_protobuf/grpc_basic/services"
	"google.golang.org/grpc"
)

const (
	NetWrotk string = "tcp"
	Address  string = ":12345"
)

func main() {

	conn, err := net.Listen(NetWrotk, Address)
	if err != nil {
		log.Println(err)
		return
	}
	rpcs := grpc.NewServer()
	protos.RegisterBidireactionalServer(rpcs, &services.Bidirection{})
	protos.RegisterClientSideServer(rpcs, &services.ClientSideService{})
	protos.RegisterEchoServiceServer(rpcs, &services.EchoService{})
	protos.RegisterServerSideServer(rpcs, &services.ServerSideService{})
	defer conn.Close()
	defer rpcs.GracefulStop()
	err = rpcs.Serve(conn)
	if err != nil {
		log.Println("grpcServer.Server err :", err)
	}
}
