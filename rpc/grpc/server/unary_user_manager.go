package server

import (
	"context"
	"rpctest/grpc/protocol/pb"
	"sync"

	"github.com/google/uuid"
)

type UserManagerServer struct {
	pb.UnimplementedUserMangerServiceServer
	users map[string]*pb.User
	mu    sync.Mutex
}

// NewUserManagerServer is a constructor for UserManagerServer
func NewUserManagerServer() *UserManagerServer {
	return &UserManagerServer{
		users: make(map[string]*pb.User),
	}
}

func (s *UserManagerServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	uuid := uuid.New().String()

	user := &pb.User{
		Id:   uuid,
		Name: req.GetName(),
		Age:  req.GetAge(),
	}

	s.users[uuid] = user

	return &pb.CreateUserResponse{User: user}, nil
}

func (s *UserManagerServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	user, ok := s.users[req.GetId()]
	if !ok {
		return nil, nil
	}
	return &pb.GetUserResponse{User: user}, nil
}

func (s *UserManagerServer) ListUsers(ctx context.Context, req *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	users := make([]*pb.User, 0, len(s.users))
	for _, user := range s.users {
		users = append(users, user)
	}
	return &pb.ListUsersResponse{Users: users}, nil
}
