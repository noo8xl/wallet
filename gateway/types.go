package gateway

import (
	"context"
	pb "wallet/api"
)

type WalletService interface {
	CreateWallet(context.Context) error
	CreateAddress(context.Context, *pb.CreateAddressRequest) error
	GetCoinBalance(context.Context, *pb.GetCoinBalanceRequest) error
	SendSingleTransaction(context.Context, *pb.SendSingleTsxRequest) error
	SendMultTransaction(context.Context, *pb.SendMultiplyTsxRequest) error
}
