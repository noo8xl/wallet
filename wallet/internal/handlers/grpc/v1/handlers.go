package grpc

import (
	"context"
	"errors"
	"log"
	pb "wallet/gen/wallet"
)

func (svc *Handlers) CreateWallet(ctx context.Context, payload *pb.CreateWalletRequest) (*pb.WalletList, error) {
	log.Printf("New CreateWallet request received! Wallet %v", payload)
	wt := svc.CreateWalletList(payload.CustomerId)
	return wt, nil
}

func (svc *Handlers) CreatePermanentAddress(c context.Context, payload *pb.CreateAddressRequest) (*pb.WalletItem, error) {
	log.Printf("New CreatePermanentAddress request received! payload -> %v", payload)
	wt := svc.DefineBlockchainAndCreatePermanentWallet(payload.CoinName, payload.CustomerId)
	log.Printf("wallet is -> %v", wt)
	if wt == nil {
		return nil, errors.New("already exists")
	}
	return wt, nil
}

func (svc *Handlers) CreateOneTimeAddress(c context.Context, payload *pb.CreateAddressRequest) (*pb.WalletItem, error) {
	log.Printf("New CreateOneTimeAddress request received! payload -> %v", payload)
	wt := svc.DefineBlockchainAndCreateOneTimeAddress(payload.CoinName, payload.CustomerId)
	return wt, nil
}

func (svc *Handlers) GetCustomerBalance(ctx context.Context, payload *pb.GetCustomerBalanceRequest) (*pb.CustomerBalance, error) {
	log.Printf("New GetCustomerBalance request received! GetCustomerBalance %v", payload)
	balance := svc.DefineBlockchainAndGetCustomerBalance(payload.CurrencyName, payload.CustomerId)
	return balance, nil
}

func (svc *Handlers) GetCoinBalance(ctx context.Context, payload *pb.GetCoinBalanceRequest) (*pb.CoinBalance, error) {
	log.Printf("New GetCoinBalance request received! GetCoinBalance %v", payload)
	balance := svc.DefineBlockchainAndGetCoinBalance(payload.CoinName, payload.Address)
	return balance, nil
}

func (svc *Handlers) SendSingleTransaction(ctx context.Context, payload *pb.SendSingleTsxRequest) (*pb.TransactionHash, error) {
	log.Printf("New SendSingleTransaction request received! SendSingleTransaction %v", payload)
	hash := svc.DefineBlockchainAndSendSingleTsx(payload)
	return hash, nil
}

func (svc *Handlers) SendMultiplyTransaction(ctx context.Context, payload *pb.SendMultipleTsxRequest) (*pb.TransactionHash, error) {
	log.Printf("New SendMultipleTransaction request received! SendMultipleTransaction %v", payload)
	hash := svc.DefineBlockchainAndSendMultiptleTsx(payload)
	return hash, nil
}
