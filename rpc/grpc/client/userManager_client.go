package client

import (
	"context"
	"fmt"
	"log"
	"rpctest/grpc/protocol/pb"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func StartUserManagerClient() {
	conn, err := grpc.NewClient(*address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewUserMangerServiceClient(conn)

	// Create a new user
	user1 := createUser(client, "Alice", 30)
	// List all users
	listUsers(client)
	// Get a specific user by ID
	getUser(client, user1.Id)

	// Create another user
	user2 := createUser(client, "Bob", 25)
	// List all users again
	listUsers(client)
	// Get the second user by ID
	getUser(client, user2.Id)
}

func createUser(client pb.UserMangerServiceClient, name string, age int32) *pb.User {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.CreateUserRequest{Name: name, Age: age}
	resp, err := client.CreateUser(ctx, req)
	if err != nil {
		log.Fatalf("Could not create user: %v", err)
	}

	fmt.Printf("Created user: %v\n", resp.GetUser())
	return resp.GetUser()
}

func getUser(client pb.UserMangerServiceClient, id string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.GetUserRequest{Id: id}
	resp, err := client.GetUser(ctx, req)
	if err != nil {
		log.Fatalf("Could not get user: %v", err)
	}

	fmt.Printf("Got user: %v\n", resp.GetUser())
}

func listUsers(client pb.UserMangerServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.ListUsersRequest{}
	resp, err := client.ListUsers(ctx, req)
	if err != nil {
		log.Fatalf("Could not list users: %v", err)
	}

	fmt.Println("List of users:")
	for _, user := range resp.GetUsers() {
		fmt.Printf("%v\n", user)
	}
}
