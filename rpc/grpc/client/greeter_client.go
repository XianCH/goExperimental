package client

import (
	"context"
	"flag"
	"log"
	"rpctest/grpc/protocol/pb"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName = "world"
)

func StartGreeterClient() {
	flag.Parse()
	conn, err := grpc.NewClient(*address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	defer conn.Close()

	//greeter client
	c := pb.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: *name})
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}

	log.Printf("Greeting: %s", r.Message)

	r, err = c.SayHelloAgain(ctx, &pb.HelloRequest{Name: *name})
	if err != nil {
		log.Fatalf("Could not greet again: %v", err)
	}
	log.Printf("Greeting again: %s", r.Message)

	//user manager client

}
