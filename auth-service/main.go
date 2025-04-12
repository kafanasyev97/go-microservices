package main

import (
	"log"
	"net"

	pb "github.com/kafanasyev97/go-microservices/proto/auth"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("не удалось слушать порт: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAuthServiceServer(grpcServer, &AuthServer{})

	log.Println("Auth Service запущен на порту 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("не удалось запустить сервер: %v", err)
	}
}
