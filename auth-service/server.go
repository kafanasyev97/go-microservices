package main

import (
	"context"

	pb "github.com/kafanasyev97/go-microservices/proto/auth"
)

type AuthServer struct {
	pb.UnimplementedAuthServiceServer
}

func (s *AuthServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	// простая заглушка
	return &pb.LoginResponse{
		Token: "dummy-token",
	}, nil
}
