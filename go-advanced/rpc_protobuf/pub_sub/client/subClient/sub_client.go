package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/x14n/goExperimental/go_advanced/rpc_protobuf/pub_sub/pb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:12345", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc dial error :%v", err)
	}
	defer conn.Close()

	client := pb.NewPubsubServiceClient(conn)
	stream, err := client.Subscribe(context.Background(), &pb.Request{Message: "golang"})
	if err != nil {
		log.Fatalf("client use Subscribe func error:%v", err)
	}

	for {
		reply, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("stream Recv error:%v", err)
		}
		fmt.Println(reply.GetMessage())
	}
}
