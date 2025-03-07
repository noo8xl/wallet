// Package theopennetwork -> is all of ton blockchain interacting
package theopennetwork

import (
	"context"
	"fmt"
	"math/big"
	"time"
	"wallet/database"
	"wallet/lib/cache"
	"wallet/lib/exceptions"
	"wallet/lib/helpers"
	"wallet/lib/models"

	pb "wallet/api"

	"github.com/xssnick/tonutils-go/address"
	"github.com/xssnick/tonutils-go/liteclient"
	"github.com/xssnick/tonutils-go/ton"
	"github.com/xssnick/tonutils-go/ton/wallet"
)

// https://github.com/xssnick/tonutils-go#Wallet
// -> the doc is here <-

type Service struct {
	client *ton.APIClient
	db     *database.DatabaseService
	store  *cache.Store
}

func InitTonService() *Service {
	client := initTonAPIConnection()
	db := database.InitDbService()
	s := cache.InitNewStore()
	return &Service{
		client: client,
		db:     db,
		store:  s,
	}
}

func (s *Service) CreatePermanentWallet(userId int64) *pb.WalletItem {

	existedAddress := s.db.IsWalletExists(userId, "ton")
	if !existedAddress {
		return s.generateAddress(userId, 0)
	}
	exceptions.HandleAnHttpExceprion()
	return nil
}

func (s *Service) CreateOneTimeddress(userId int64) *pb.WalletItem {
	return s.generateAddress(userId, 1)
}

// GetTonBalanceByAddress -> get balance value by coin address
func (s *Service) GetBalanceByAddress(a string) *big.Float {

	result, err := s.store.GetAKey(a)
	if val := helpers.BalanceFromStoreFormatter(result, err); val != nil {
		return val
	}

	curBalance := new(big.Float)
	ctx := context.Background()
	addr := address.MustParseAddr(a)

	// we need fresh block info to run get methods
	blcn, err := s.client.CurrentMasterchainInfo(ctx)
	if err != nil {
		exceptions.HandleAnException("<ton CurrentMasterchainInfo> got an err: " + err.Error())
	}

	// we use WaitForBlock to make sure block is ready,
	// it is optional but escapes us from liteserver block not ready errors
	acc, err := s.client.WaitForBlock(blcn.SeqNo).GetAccount(ctx, blcn, addr)
	if err != nil {
		exceptions.HandleAnException("<ton GetAccount> got an err: " + err.Error())
	}

	fmt.Println("acc info =>")
	helpers.PrintPretty(blcn)

	fmt.Printf("Is active: %v\n", acc.IsActive)
	if !acc.IsActive {
		exceptions.HandleAnException("<ton account.IsActive> got an err: " + err.Error())
	}

	fmt.Printf("Status: %s\n", acc.State.Status)

	curBalance.SetString(acc.State.Balance.String())
	// fmt.Println(" curBalance --> ", curBalance)
	s.store.SetAKey(a, curBalance.String())
	return curBalance
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
