package gateway

import (
	"context"
	"log"
	pb "wallet/api"
	"wallet/crypto-lib/common"

	"google.golang.org/grpc"
)

type grpcHandler struct {
	pb.UnimplementedWalletServiceServer
	service WalletService
	common.ServiceType
}

func NewGRPCHandler(grpcServer *grpc.Server, service WalletService) {
	handler := &grpcHandler{
		service: service,
	}
	pb.RegisterWalletServiceServer(grpcServer, handler)
}

func (h *grpcHandler) CreateWallet(ctx context.Context, payload *pb.CreateWalletRequest) (*pb.WalletList, error) {
	log.Printf("New request received! Wallet %v", payload)
	wt := h.CreateWalletList(payload.CustomerId)
	return wt, nil
}

func (h *grpcHandler) CreateAddress(ctx context.Context, payload *pb.CreateAddressRequest) (*pb.WalletItem, error) {
	log.Printf("New request received! Address %v", payload)
	wt := h.DefineBlockchainAndCreateWallet(payload.CoinName, payload.CustomerId)
	return wt, nil
}

func (h *grpcHandler) GetCustomerBalance(ctx context.Context, payload *pb.GetCustomerBalanceRequest) (*pb.CustomerBalance, error) {
	log.Printf("New request received! GetCustomerBalance %v", payload)
	balance := h.DefineBlockchainAndGetCustomerBalance(payload.CurrencyName, payload.CustomerId)
	return balance, nil
}

func (h *grpcHandler) GetCoinBalance(ctx context.Context, payload *pb.GetCoinBalanceRequest) (*pb.CoinBalance, error) {
	log.Printf("New request received! GetCoinBalance %v", payload)
	balance := h.DefineBlockchainAndGetCoinBalance(payload.CoinName, payload.Address)
	return balance, nil
}

func (h *grpcHandler) SendSingleTransaction(ctx context.Context, payload *pb.SendSingleTsxRequest) (*pb.TransactionHash, error) {
	log.Printf("New request received! SendSingleTransaction %v", payload)
	hash := h.DefineBlockchainAndSendSingleTsx(payload)
	return hash, nil
}

func (h *grpcHandler) SendMultiplyTransaction(ctx context.Context, payload *pb.SendMultiplyTsxRequest) (*pb.TransactionHash, error) {
	log.Printf("New request received! SendMultiplyTransaction %v", payload)
	hash := h.DefineBlockchainAndSendMultiptlyTsx(payload)
	return hash, nil
}
