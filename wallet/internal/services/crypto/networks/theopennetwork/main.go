// Package theopennetwork -> is all of ton blockchain interacting
package theopennetwork

import (
	"context"
	"fmt"
	"time"
	"wallet/pkg/exceptions"
	"wallet/wallet/internal/repository/mysql"
	"wallet/wallet/internal/repository/redis"
	models "wallet/wallet/pkg/models/wallet"

	pb "wallet/gen/wallet"

	"github.com/xssnick/tonutils-go/liteclient"
	"github.com/xssnick/tonutils-go/ton"
	"github.com/xssnick/tonutils-go/ton/wallet"
)

// https://github.com/xssnick/tonutils-go#Wallet
// -> the doc is here <-

type Service struct {
	client *ton.APIClient
	db     *mysql.DatabaseService
	store  *redis.Store
}

func InitService() *Service {
	client := initTonAPIConnection()
	db := mysql.InitDbService()
	s := redis.InitNewStore()
	return &Service{
		client: client,
		db:     db,
		store:  s,
	}
}

// ===========================================================================================//
// ============================ init the blockchain connection ===============================//
// ===========================================================================================//

func (s *Service) generateAddress(userId int64, opt byte) *pb.WalletItem {
	words := wallet.NewSeed()
	w, err := wallet.FromSeed(s.client, words, wallet.V4R2)
	if err != nil {
		exceptions.HandleAnException("<ton GenAddrKeychain> got an err: " + err.Error())
	}

	wt := &models.TonWallet{
		Address:    w.WalletAddress().String(),
		AddrType:   w.Address().Type(),
		PrivateKey: w.PrivateKey(),
		BitsLen:    w.Address().BitsLen(),
		CreatedAt:  time.Now().UnixMilli(),
		UpdatedAt:  time.Now().UnixMilli(),
		CustomerId: userId,
	}

	// -> save wallet to db!
	switch opt {
	case 0:
		if err := s.db.InsertTonWalletToPermanent(wt); err != nil {
			exceptions.HandleAnException("<InsertTonWalletToPermanent> got an error: " + err.Error())
		}
	case 1:
		if err := s.db.InsertTonWalletToOneTimeAddresses(wt); err != nil {
			exceptions.HandleAnException("<InsertTonWalletToOneTimeAddresses> got an error: " + err.Error())
		}
	default:
		exceptions.HandleAnException(fmt.Sprintf("Unknown opt value %d", opt))
	}

	return &pb.WalletItem{CoinName: "ton", Address: wt.Address}
}

// InitTonAPIConnection -> create ton blockchain connection
func initTonAPIConnection() *ton.APIClient {

	// configUrl := "https://ton.org/global.config.json" // -> Mainnet
	configUrl := "https://ton-blockchain.github.io/testnet-global.config.json" // -> Testnet

	client := liteclient.NewConnectionPool()
	ctx := client.StickyContext(context.Background())
	err := client.AddConnectionsFromConfigUrl(ctx, configUrl)
	if err != nil {
		exceptions.HandleAnException("Init <ton> blockchain got an error: " + err.Error())
	}
	return ton.NewAPIClient(client)
	// api = api.WithRetry() // if you want automatic retries with failover to another node
}
