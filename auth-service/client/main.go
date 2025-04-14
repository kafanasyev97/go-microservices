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
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// Register
	regResp, err := client.Register(ctx, &pb.RegisterRequest{
		Username: "john",
		Password: "pass123",
	})
	if err != nil {
		log.Fatalf("ошибка регистрации: %v", err)
	}
	log.Printf("Зарегистрирован пользователь с ID: %s\n", regResp.UserId)

	// Login
	loginResp, err := client.Login(ctx, &pb.LoginRequest{
		Username: "john",
		Password: "pass123",
	})
	if err != nil {
		log.Fatalf("ошибка логина: %v", err)
	}
	log.Printf("Получен токен: %s\n", loginResp.Token)

	// ValidateToken
	validateResp, err := client.ValidateToken(ctx, &pb.ValidateTokenRequest{
		Token: loginResp.Token,
	})
	if err != nil {
		log.Fatalf("ошибка проверки токена: %v", err)
	}
	log.Printf("Токен валиден: %v, user_id: %s\n", validateResp.Valid, validateResp.UserId)
}
