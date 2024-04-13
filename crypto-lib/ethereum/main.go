// Package ethereumwallet -> is eth network interacting
package ethereum

import (
	"fmt"
	"math"
	"math/big"
	"wallet-cli/lib/config"
	"wallet-cli/lib/helpers"

	"github.com/blockcypher/gobcy/v2"
)

// get api token here
var apiToken string = config.GetEthereumAPIKey()
var bc gobcy.API

// ----------------------------------------------------------------

// CreateWallet is in charge of creating a new root wallet
func CreateWallet(userID string) bool {

	var err error
	var w gobcy.Wallet
	var addressKeys gobcy.AddrKeychain

	initBlockchain("eth")
	addressKeys, err = bc.GenAddrKeychain()
	if err != nil {
		return false
	}

	fmt.Println("generated address is\n->", addressKeys)
	// normal wallet
	w, err = bc.CreateWallet(gobcy.Wallet{
		Name:      "test",
		Addresses: []string{addressKeys.Address},
	})
	if err != nil {
		return false
	} else {

		// -> save wallet were !

		fmt.Printf("Normal Wallet:%+v\n", w)
		return true
	}

}

// GetEthereumAddressBalance -> get balance by address
func GetEthereumAddressBalance(addr string) *big.Float {
	var addressData gobcy.Addr
	var err error
	var currentBalance *big.Float
	initBlockchain("eth")

	addressData, err = bc.GetAddrBal(addr, nil)
	if err != nil {
		return currentBalance
	}

	fmt.Println("addressData -> ")
	helpers.PrintPretty(addressData)

	currentBalance = new(big.Float)
	currentBalance.SetString(addressData.Balance.String())
	ethValue := new(big.Float).Quo(currentBalance, big.NewFloat(math.Pow10(18)))

	fmt.Println("balance -> ", ethValue)
	return ethValue
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
		bc.Coin = c       //options: "btc","bcy","ltc","doge","eth"
		bc.Chain = "main" //depending on coin: "main","test3","test"

		fmt.Println("blockchain  ->", bc)
	}

}

// func saveWallet(gobcy.AddrKeychain) bool {

// 	return true
// }
