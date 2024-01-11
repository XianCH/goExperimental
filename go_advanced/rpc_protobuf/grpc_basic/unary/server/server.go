package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/x14n/goExperimental/go_advanced/rpc_protobuf/grpc_basic/protos"
	"google.golang.org/grpc"
)

type echoService struct {
	protos.UnimplementedEchoServiceServer
}

// GetUnaryEcho(context.Context, *EchoRequest) (*EchoResponse, error)

func (es *echoService) GetUnaryEcho(ctx context.Context, req *protos.EchoRequest) (*protos.EchoResponse, error) {
	res := "revies" + req.GetReq()
	fmt.Println(res)

	return &protos.EchoResponse{Res: res}, nil
}

func main() {
	rpcs := grpc.NewServer()
	protos.RegisterEchoServiceServer(rpcs, new(echoService))

	conn, err := net.Listen("tcp", ":12345")
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()
	rpcs.Serve(conn)
}
