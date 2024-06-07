package server

import (
	"flag"
	"fmt"
	"log"
	"net"

	"rpctest/grpc/protocol/pb"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func StartRpcServer() {
	flag.Parse()
	log.Printf("Server started on port %d", *port)
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create a new gRPC server
	s := grpc.NewServer()
	// Register the server with the gRPC server

	//register greeter server
	// pb.RegisterGreeterServer(s, &server{})

	//register userManager server
	userserver := NewUserManagerServer()
	pb.RegisterUserMangerServiceServer(s, userserver)

	log.Printf("server listening at %v", l.Addr())
	if err := s.Serve(l); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
