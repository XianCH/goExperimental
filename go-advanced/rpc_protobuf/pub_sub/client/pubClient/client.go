package main

import (
	"context"
	"log"

	"github.com/x14n/goExperimental/go_advanced/rpc_protobuf/pub_sub/pb"
	"google.golang.org/grpc"
)

const (
	Address string = "127.0.0.1:12345"
	NetWrok string = "tcp"
)

func main() {
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc dial error:%v", err)
	}
	defer conn.Close()

	client := pb.NewPubsubServiceClient(conn)
	_, err = client.Publish(
		context.Background(), &pb.Request{Message: "golang: Hello Go!"})
	if err != nil {
		log.Fatalf("client publish func error :%v", err)
	}
	_, err = client.Publish(
		context.Background(), &pb.Request{Message: "all: Hello World!"})
	if err != nil {
		log.Fatalf("client publish func error :%v", err)
	}
}
