// Package bitcoinwallet -> is all of btc network interacting
package bitcoin

import (
	"math/big"
	"time"
	"wallet-cli/config"
	"wallet-cli/database"
	"wallet-cli/lib/exceptions"
	"wallet-cli/lib/models"

	"github.com/blockcypher/gobcy/v2"
)

// CreateWallet is in charge of creating a new root wallet
func CreateWallet(userId *string) *models.WalletListItem {

	bc := initBlockchain("btc")
	stamp := time.Now().UnixMilli()

	addressKeys, err := bc.GenAddrKeychain()
	if err != nil {
		exceptions.HandleAnException("<btc GenAddrKeychain> got an err: " + err.Error())
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
			UserId:          *userId,
			// PubKeys:         addressKeys.PubKeys,
		}

		// -> save wallet to main db <-
		if err := database.InsertBtcWallet(&wt); err != nil {
			exceptions.HandleAnException("<Database insertion> got an error: " + err.Error())
		}
	}

	return &models.WalletListItem{CoinName: "btc", Address: addressKeys.Address}
}

func CreateOneTimeBitcoinAddress(userID string) (string, error) {
	bc := initBlockchain("btc")

	addressKeys, err := bc.GenAddrKeychain()
	if err != nil {
		exceptions.HandleAnException("<btc GenAddrKeychain> got an err: " + err.Error())
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
	bc := initBlockchain("btc")

	addressData, err := bc.GetAddrBal(address, nil)
	if err != nil {
		exceptions.HandleAnException("<btc GetAddrBal> got an error: " + err.Error())
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

func initBlockchain(c string) *gobcy.API {

	bc := new(gobcy.API)
	apiToken := config.GetBitcoinAPIKey()

	if apiToken == "" {
		exceptions.HandleAnException("Init <btc> blockchain got an error. Empty API token.")
	} else {

		bc.Token = apiToken
		bc.Coin = c        // options: "btc","bcy","ltc","doge","eth"
		bc.Chain = "test3" // depend on coin: "main","test3","test"\
		// bc.Chain = "main"
	}

	return bc
}
