package main

import (
	"log"
	"net/http"
	"wallet/lib/exceptions"

	pb "wallet/api"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	grpcServiceAddress = "127.0.0.1:20002"
	httpServerAddress  = "127.0.0.1:20003"
)

func main() {

	conn, err := grpc.NewClient(grpcServiceAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		exceptions.HandleAnException("Got an error at <create grpc client>: " + err.Error())
	}
	defer conn.Close()

	log.Println("Server was started at ", grpcServiceAddress)

	client := pb.NewWalletServiceClient(conn)

	mux := http.NewServeMux()
	handler := NewHandler(client)
	handler.registerRoutes(mux)

	err = http.ListenAndServe(httpServerAddress, mux)
	if err != nil {
		log.Fatal("Failed to start http server")
	}

	log.Printf("server is running successfuly on %s", httpServerAddress)

}

// https://mahdidarabi.medium.com/create-raw-bitcoin-transaction-and-sign-it-with-golang-96b5e10c30aa
// https://goethereumbook.org/wallet-generate/

// https://hackernoon.com/how-to-create-a-bitcoin-hd-wallet-with-golang-and-grpc-part-l-u51d3wwm

// https://github.com/xssnick/tonutils-go
