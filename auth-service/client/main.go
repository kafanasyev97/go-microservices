package main

import (
	"context"
	"log"
	"time"

	pb "github.com/kafanasyev97/go-microservices/proto/auth"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("не удалось подключиться: %v", err)
	}
	defer conn.Close()

	client := pb.NewAuthServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.Login(ctx, &pb.LoginRequest{
		Username: "testuser",
		Password: "secret",
	})
	if err != nil {
		log.Fatalf("ошибка при логине: %v", err)
	}

	log.Printf("Получен токен: %s", resp.Token)
}
