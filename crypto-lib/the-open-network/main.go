// Package theopennetwork -> is all of ton blockchain interacting
package theopennetwork

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"time"
	"wallet-cli/database"
	"wallet-cli/lib/helpers"
	"wallet-cli/lib/models"

	"github.com/xssnick/tonutils-go/address"
	"github.com/xssnick/tonutils-go/liteclient"
	"github.com/xssnick/tonutils-go/tlb"
	"github.com/xssnick/tonutils-go/ton"
	"github.com/xssnick/tonutils-go/ton/wallet"
)

// ton api client instance
// var tonAPI *ton.APIClient
var ctx context.Context

// https://github.com/xssnick/tonutils-go#Wallet
// -> the doc is here <-

func CreateWallet(userId string) string {
	var err error
	var words []string
	var w *wallet.Wallet
	tonAPI := initTonAPIConnection()
	stamp := time.Now().UnixMilli()

	words = wallet.NewSeed()
	w, err = wallet.FromSeed(tonAPI, words, wallet.V4R2)
	if err != nil {
		return ""
	} else {
		wt := models.TonWallet{
			Address:    w.WalletAddress().String(),
			AddrType:   w.Address().Type(),
			PrivateKey: w.PrivateKey(),
			BitsLen:    w.Address().BitsLen(),
			CreatedAt:  stamp,
			UpdatedAt:  stamp,
			UserId:     userId,
		}
		// -> save wallet were to main db!
		if err := database.InsertTonWallet(wt); err != nil {
			log.Println("database insertion error", err)
			os.Exit(1)
		}
	}

	log.Println("wallet:\n", w.WalletAddress().String())
	return w.WalletAddress().String()
}

// GetTonBalanceByAddress -> get balance value by coin address
func GetTonBalanceByAddress(a string) *big.Float {

	curBalance := new(big.Float)
	var addr *address.Address
	var blcn *ton.BlockIDExt
	var acc *tlb.Account
	var err error
	ctx = context.Background()

	tonAPI := initTonAPIConnection()
	addr = address.MustParseAddr(a)

	// we need fresh block info to run get methods
	blcn, err = tonAPI.CurrentMasterchainInfo(ctx)
	if err != nil {
		return nil
	}

	// we use WaitForBlock to make sure block is ready,
	// it is optional but escapes us from liteserver block not ready errors
	acc, err = tonAPI.WaitForBlock(blcn.SeqNo).GetAccount(ctx, blcn, addr)
	if err != nil {
		return nil
	}

	fmt.Println("acc info =>")
	helpers.PrintPretty(blcn)

	fmt.Printf("Is active: %v\n", acc.IsActive)
	if !acc.IsActive {
		return nil
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
func initTonAPIConnection() *ton.APIClient {

	// configUrl := "https://ton.org/global.config.json" // -> Mainnet
	configUrl := "https://ton-blockchain.github.io/testnet-global.config.json" // -> Testnet

	client := liteclient.NewConnectionPool()
	ctx = client.StickyContext(context.Background())
	err := client.AddConnectionsFromConfigUrl(ctx, configUrl)
	if err != nil {
		fmt.Println("ton init error.")
		os.Exit(1)
	}
	return ton.NewAPIClient(client)
	// api = api.WithRetry() // if you want automatic retries with failover to another node
}
