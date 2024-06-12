package client

import (
	"bufio"
	"context"
	"log"
	"os"
	"rpctest/grpc/protocol/pb"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func StartChatClient() {

	conn, err := grpc.NewClient(*address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewChatServiceClient(conn)

	// ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	// defer cancel()

	stream, err := c.Chat(context.Background())
	if err != nil {
		log.Fatalf("Could not chat: %v", err)
	}

	scanner := bufio.NewScanner(os.Stdin)
	log.Print("Enter message: ")
	scanner.Scan()
	username := scanner.Text()

	//start a go routine to receive messages
	go func() {
		for {
			in, err := stream.Recv()
			if err != nil {
				log.Fatalf("Failed to receive message: %v", err)
			}
			log.Printf("Received message from: %s: %s", in.User, in.Message)
		}
	}()

	for {
		log.Print("Enter message: ")
		scanner.Scan()
		test := scanner.Text()

		if err := stream.Send(&pb.ChatMessage{User: username, Message: test}); err != nil {
			log.Fatalf("Failed to send message: %v", err)
		}
		time.Sleep(time.Second)
	}

}
