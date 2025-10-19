package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	handlers "wallet/wallet/internal/handlers/grpc/v1"
)

func main() {
	address := "127.0.0.1:20002"
	server := grpc.NewServer()

	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer lis.Close()

	handlers.NewGRPCHandler(server)

	log.Printf("grpc server started on %s", address)

	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to start grpc server: %v", err)
	}

}
