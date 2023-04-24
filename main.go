package main

import (
	"log"
	"net"
	"os"

	"github.com/MuhAndriJP/gateway-service.git/routes"
	"google.golang.org/grpc"
)

func main() {
	e := routes.New()

	lis, err := net.Listen("tcp", ":"+os.Getenv("port_grpc"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	} else {
		log.Println("Connect GRPC to port " + os.Getenv("port_grpc"))
	}

	e.Logger.Fatal(e.Start(":8080"))
	grpcServer := grpc.NewServer()
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
