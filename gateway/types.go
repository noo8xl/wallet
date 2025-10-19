package main

import (
	"context"
	pb "wallet/gen/wallet"
)

type WalletService interface {
	CreatePermanentAddress(context.Context) error
	CreateOneTimeAddress(context.Context, *pb.CreateAddressRequest) error
	GetCoinBalance(context.Context, *pb.GetCoinBalanceRequest) error
	SendSingleTransaction(context.Context, *pb.SendSingleTsxRequest) error
	SendMultTransaction(context.Context, *pb.SendMultipleTsxRequest) error
}
