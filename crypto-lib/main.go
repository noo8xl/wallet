package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

var (
	grpcServiceAddress = "127.0.0.1:20002"
)

func main() {
	server := grpc.NewServer()

	lis, err := net.Listen("tcp", grpcServiceAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer lis.Close()

	NewGRPCHandler(server)

	log.Printf("grpc server started on %s", grpcServiceAddress)

	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to start grpc server: %v", err)
	}

}
