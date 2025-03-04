package main

import (
	"context"
	"errors"
	"log"
	pb "wallet/api"
	"wallet/crypto-lib/common"

	"google.golang.org/grpc"
)

type service struct {
	pb.UnimplementedWalletServiceServer
	common.Service
}

func NewService() *service {
	s := common.InitService()
	return &service{
		Service: *s,
	}
}

func NewGRPCHandler(grpcServer *grpc.Server) {
	handler := NewService()
	pb.RegisterWalletServiceServer(grpcServer, handler)
}

func (svc *service) CreateWallet(ctx context.Context, payload *pb.CreateWalletRequest) (*pb.WalletList, error) {
	log.Printf("New CreateWallet request received! Wallet %v", payload)
	wt := svc.CreateWalletList(payload.CustomerId)
	return wt, nil
}

func (svc *service) CreatePermanentAddress(c context.Context, payload *pb.CreateAddressRequest) (*pb.WalletItem, error) {
	log.Printf("New CreatePermanentAddress request received! payload -> %v", payload)
	wt := svc.DefineBlockchainAndCreatePermanentWallet(payload.CoinName, payload.CustomerId)
	log.Printf("wallet is -> %v", wt)
	if wt == nil {
		return nil, errors.New("already exists")
	}
	return wt, nil
}

func (svc *service) CreateOneTimeAddress(c context.Context, payload *pb.CreateAddressRequest) (*pb.WalletItem, error) {
	log.Printf("New CreateOneTimeAddress request received! payload -> %v", payload)
	wt := svc.DefineBlockchainAndCreateOneTimeAddress(payload.CoinName, payload.CustomerId)
	return wt, nil
}

func (svc *service) GetCustomerBalance(ctx context.Context, payload *pb.GetCustomerBalanceRequest) (*pb.CustomerBalance, error) {
	log.Printf("New GetCustomerBalance request received! GetCustomerBalance %v", payload)
	balance := svc.DefineBlockchainAndGetCustomerBalance(payload.CurrencyName, payload.CustomerId)
	return balance, nil
}

func (svc *service) GetCoinBalance(ctx context.Context, payload *pb.GetCoinBalanceRequest) (*pb.CoinBalance, error) {
	log.Printf("New GetCoinBalance request received! GetCoinBalance %v", payload)
	balance := svc.DefineBlockchainAndGetCoinBalance(payload.CoinName, payload.Address)
	return balance, nil
}

func (svc *service) SendSingleTransaction(ctx context.Context, payload *pb.SendSingleTsxRequest) (*pb.TransactionHash, error) {
	log.Printf("New SendSingleTransaction request received! SendSingleTransaction %v", payload)
	hash := svc.DefineBlockchainAndSendSingleTsx(payload)
	return hash, nil
}

func (svc *service) SendMultiplyTransaction(ctx context.Context, payload *pb.SendMultiplyTsxRequest) (*pb.TransactionHash, error) {
	log.Printf("New SendMultiplyTransaction request received! SendMultiplyTransaction %v", payload)
	hash := svc.DefineBlockchainAndSendMultiptlyTsx(payload)
	return hash, nil
}
