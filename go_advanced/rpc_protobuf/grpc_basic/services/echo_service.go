package services

import (
	"context"
	"fmt"

	"github.com/x14n/goExperimental/go_advanced/rpc_protobuf/grpc_basic/protos"
)

type EchoService struct {
	protos.UnimplementedEchoServiceServer
}

// GetUnaryEcho(context.Context, *EchoRequest) (*EchoResponse, error)
func (e *EchoService) GetUnaryEcho(ctx context.Context, req *protos.EchoRequest) (*protos.EchoResponse, error) {
	res := "revice:" + req.GetReq()
	fmt.Println(res)
	return &protos.EchoResponse{Res: res}, nil
}
