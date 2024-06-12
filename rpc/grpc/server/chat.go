package server

import (
	"log"
	"rpctest/grpc/protocol/pb"
	"sync"
)

type ChatServer struct {
	pb.UnimplementedChatServiceServer
	mu      sync.Mutex
	clients map[pb.ChatService_ChatServer]struct{}
}

func NewChatServer() *ChatServer {
	return &ChatServer{
		clients: make(map[pb.ChatService_ChatServer]struct{}),
	}
}

func (s *ChatServer) Chat(stream pb.ChatService_ChatServer) error {
	s.mu.Lock()
	s.clients[stream] = struct{}{}
	s.mu.Unlock()

	defer func() {
		s.mu.Lock()
		delete(s.clients, stream)
		s.mu.Unlock()
	}()

	for {
		msg, err := stream.Recv()
		if err != nil {
			log.Printf("Failed to receive message: %v", err)
			return err
		}
		log.Printf("Received message from: %s : %s", msg.User, msg.Message)
		s.Broadcast(msg)
	}
}

func (s *ChatServer) Broadcast(message *pb.ChatMessage) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for stream := range s.clients {
		if err := stream.Send(message); err != nil {
			log.Printf("Failed to send message to client: %v", err)
		}
	}
}
