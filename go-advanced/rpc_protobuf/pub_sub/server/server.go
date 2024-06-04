package main

import (
	"fmt"
	"log"
	"net"

	"github.com/x14n/goExperimental/go_advanced/rpc_protobuf/pub_sub/pb"
	"github.com/x14n/goExperimental/go_advanced/rpc_protobuf/pub_sub/service"
	"google.golang.org/grpc"
)

const (
	address string = ":12345"
	netWrok string = "tcp"
)

func main() {
	conn, err := net.Listen(netWrok, address)
	if err != nil {
		log.Fatalf("server start error:%v", err)
	}
	fmt.Println("server start:", address)
	defer conn.Close()
	rpcs := grpc.NewServer()
	pb.RegisterPubsubServiceServer(rpcs, &service.PubsubService{})
	defer rpcs.GracefulStop()
	err = rpcs.Serve(conn)
	if err != nil {
		log.Println("grpcServer.Server err :", err)
	}
}
