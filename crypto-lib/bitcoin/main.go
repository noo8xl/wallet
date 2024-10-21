// Package bitcoinwallet -> is all of btc network interacting
package bitcoin

import (
	"fmt"
	"log"
	"math/big"
	"os"
	"time"
	"wallet-cli/config"
	"wallet-cli/database"
	"wallet-cli/lib/models"

	"github.com/blockcypher/gobcy/v2"
)

var bc gobcy.API
var apiToken = config.GetBitcoinAPIKey()

// CreateWallet is in charge of creating a new root wallet
func CreateWallet(userID string) string {

	initBlockchain("btc")
	stamp := time.Now().UnixMilli()

	addressKeys, err := bc.GenAddrKeychain()
	if err != nil {
		panic(err)
	} else {
		wt := models.BtcWallet{
			Address:         addressKeys.Address,
			PrivateKey:      addressKeys.Private,
			PublicKey:       addressKeys.Public,
			Wif:             addressKeys.Wif,
			ScriptType:      addressKeys.ScriptType,
			OriginalAddress: addressKeys.OriginalAddress,
			OAPAddress:      addressKeys.OAPAddress,
			CreatedAt:       stamp,
			UpdatedAt:       stamp,
			UserId:          userID,
			// PubKeys:         addressKeys.PubKeys,
		}

		// -> save wallet to main db <-
		if err := database.InsertBtcWallet(wt); err != nil {
			log.Println("database insertion error", err)
			os.Exit(1)
		}
	}

	// returt btc address
	fmt.Println("generated btc address is\n->", addressKeys)
	return addressKeys.Address
}

func CreateOneTimeBitcoinAddress(userID string) (string, error) {
	var err error
	var addressKeys gobcy.AddrKeychain

	initBlockchain("btc")

	addressKeys, err = bc.GenAddrKeychain()
	if err != nil {
		return "", err
	}

	return addressKeys.Address, nil
}

// GetBitcoinAddressBalance -> get balance by address
func GetBitcoinAddressBalance(address string) *big.Float {

	// log.Println("adr -> ", address)
	// ###################################################
	// ######## DOESN"T WORK IN THE TEST-NET! ############
	// ###################################################

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
		bc.Coin = c // options: "btc","bcy","ltc","doge","eth"
		// bc.Chain = "test3" // depend on coin: "main","test3","test"\
		bc.Chain = "test3"

		fmt.Println("blockchain is->", bc.Chain)
	}
}
