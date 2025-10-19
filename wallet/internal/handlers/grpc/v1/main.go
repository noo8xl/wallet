package grpc

import (
	pb "wallet/gen/wallet"
	"wallet/wallet/internal/services/crypto"

	"google.golang.org/grpc"
)

type Handlers struct {
	pb.UnimplementedWalletServiceServer
	crypto.CryptoService
}

func NewGRPCHandler(grpcServer *grpc.Server) {
	handler := newService()
	pb.RegisterWalletServiceServer(grpcServer, handler)
}

func newService() *Handlers {
	s := crypto.InitService()
	return &Handlers{
		CryptoService: *s,
	}
}
