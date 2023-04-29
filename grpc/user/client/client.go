package client

import (
	"context"
	"log"
	"os"

	pb "github.com/MuhAndriJP/gateway-service.git/grpc/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
}

func (c *Client) RegisterUser(ctx context.Context, req *pb.RegisterUserRequest) (*pb.NoResponse, error) {
	conn, err := grpc.Dial(os.Getenv("USER_GRPC"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewUserClient(conn)

	return client.RegisterUser(ctx, req)
}

func (c *Client) LoginUser(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	conn, err := grpc.Dial(os.Getenv("USER_GRPC"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewUserClient(conn)

	return client.LoginUser(ctx, req)
}

func (c *Client) GetUserByEmail(ctx context.Context, req *pb.GetUserByEmailRequest) (*pb.GetUserResponse, error) {
	conn, err := grpc.Dial(os.Getenv("USER_GRPC"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewUserClient(conn)

	return client.GetUserByEmail(ctx, req)
}
