// Package theopennetwork -> is all of ton blockchain interacting
package theopennetwork

import (
	"context"
	"fmt"
	"math/big"
	"time"
	"wallet/database"
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

type TONService struct {
	client *ton.APIClient
	db     database.DatabaseService
}

func init() {
	initTonAPIConnection()
}

func (s *TONService) CreateWallet(userId int64) *pb.WalletItem {

	words := wallet.NewSeed()
	w, err := wallet.FromSeed(s.client, words, wallet.V4R2)
	if err != nil {
		exceptions.HandleAnException("<ton GenAddrKeychain> got an err: " + err.Error())
	} else {
		wt := models.TonWallet{
			Address:    w.WalletAddress().String(),
			AddrType:   w.Address().Type(),
			PrivateKey: w.PrivateKey(),
			BitsLen:    w.Address().BitsLen(),
			CreatedAt:  time.Now().UnixMilli(),
			UpdatedAt:  time.Now().UnixMilli(),
			CustomerId: userId,
		}
		// -> save wallet were to db!
		if err := s.db.InsertTonWallet(&wt); err != nil {
			exceptions.HandleAnException("<Database insertion> got an error: " + err.Error())
		}
	}

	return &pb.WalletItem{CoinName: "ton", Address: w.WalletAddress().String(), Balance: 0.0}
}

// GetTonBalanceByAddress -> get balance value by coin address
func (s *TONService) GetTonBalanceByAddress(a string) *big.Float {

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
	fmt.Println(" curBalance --> ", curBalance)

	return curBalance
}

// ===========================================================================================//
// ============================ init the blockchain connection ===============================//
// ===========================================================================================//

// InitTonAPIConnection -> create ton blockchain connection
func initTonAPIConnection() *TONService {

	// configUrl := "https://ton.org/global.config.json" // -> Mainnet
	configUrl := "https://ton-blockchain.github.io/testnet-global.config.json" // -> Testnet

	client := liteclient.NewConnectionPool()
	ctx := client.StickyContext(context.Background())
	err := client.AddConnectionsFromConfigUrl(ctx, configUrl)
	if err != nil {
		exceptions.HandleAnException("Init <ton> blockchain got an error: " + err.Error())
	}
	return &TONService{client: ton.NewAPIClient(client)}
	// api = api.WithRetry() // if you want automatic retries with failover to another node
}
