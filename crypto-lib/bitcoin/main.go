// Package bitcoinwallet -> is all of btc network interacting
package bitcoin

import (
	"fmt"
	"log"
	"math/big"
	"strconv"
	"wallet-cli/lib/config"
	"wallet-cli/lib/models"

	"github.com/blockcypher/gobcy/v2"
)

var bc gobcy.API
var apiToken = config.GetBitcoinAPIKey()

// CreateWallet is in charge of creating a new root wallet
func CreateWallet(userID string) bool {
	initBlockchain("btc")
	var userW models.WalletCoinItem
	uId, _ := strconv.Atoi(userID)

	addressKeys, err := bc.GenAddrKeychain()
	if err != nil {
		return false
	} else {
		fmt.Println("generated btc address is\n->", addressKeys)
		// -> save wallet to main mysql db (in cryptomodule)! <-

		userW.UserId = int64(uId)
		userW.Address = addressKeys.Address
		userW.CoinBalance = 0
		userW.CoinName = "btc"
		userW.FiatBalance = 0
		userW.WalletId = "" // update before save to db ()

		// save cache data by userId as models.WalletCoinItem <-----
		// return cache.SetNewWalletCoinItemData(userW)
		return false
	}

	// // // normal wallet
	// w, err = bc.CreateWallet(gobcy.Wallet{
	// 	// Name:      "test",
	// 	Addresses: []string{addressKeys.Address},
	// })
	// if err != nil {
	// 	exception.ErrorHandlerException("create btc wallet", err)
	// 	return false
	// } else {
	// 	// -> save wallet were !

	// 	fmt.Println("w is ->\n", w)
	// 	return true
	// }
}

func CreateOneTimeBitcoinAddress(userID string) (string, error) {
	var err error
	var addressKeys gobcy.AddrKeychain

	initBlockchain("btc")

	addressKeys, err = bc.GenAddrKeychain()
	if err != nil {
		return "", err
	}

	// cache.SetNewWalletCoinItemData(addressKeys)

	return addressKeys.Address, nil
}

// GetBitcoinAddressBalance -> get balance by address
func GetBitcoinAddressBalance(address string) *big.Float {

	log.Println("adr -> ", address)
	// ###############################################
	// ######## DOESN"T WORK AT TEST NET! ############
	// ###############################################

	// balance value from API will be receive in satoshi value
	// to calculate it in btc value -> should multiply on 0.00_000_001
	// https://buybitcoinworldwide.com/satoshi-usd/

	const satoshiPerByte float64 = 0.00000001
	var currentBalance = big.NewFloat(0)
	var addressData gobcy.Addr
	var err error
	initBlockchain("btc")

	addressData, err = bc.GetAddrBal(address, nil)
	if err != nil {
		fmt.Println("log from err -> ", err)
		return currentBalance
	}

	// fmt.Println("addressData -> ")
	// helpers.PrintPretty(addressData)
	// fmt.Println("addressData.Balance -> ", addressData.Balance.String())

	currentBalance = new(big.Float)
	currentBalance.SetPrec(30)
	currentBalance.SetString(addressData.Balance.String())

	bal := new(big.Float).Mul(currentBalance, big.NewFloat(satoshiPerByte))

	// tested with this wallet ->
	// -> bc1qf27zs0gyyfkankywsppx5dxgxm8lufgh9jhxw9
	// expected value is -> 0.00_000_598
	// but receive -> 5.98e-06

	return bal
}

// ===========================================================================================//
// ============================ init the blockchain connection ===============================//
// ===========================================================================================//

func initBlockchain(c string) {
	if apiToken == "" {
		fmt.Println("init blockchain error. empty API token.")
		return
	} else {
		bc = gobcy.API{}
		bc.Token = apiToken
		bc.Coin = c       // options: "btc","bcy","ltc","doge","eth"
		bc.Chain = "main" // depending on coin: "main","test3","test"

		fmt.Println("blockchain is->", bc.Chain)
	}
}
