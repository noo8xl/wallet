// Package ethereumwallet -> is eth network interacting
package ethereum

import (
	"fmt"
	"math"
	"math/big"
	"time"
	"wallet-cli/config"
	"wallet-cli/database"
	"wallet-cli/lib/exceptions"
	"wallet-cli/lib/models"

	"github.com/blockcypher/gobcy/v2"
)

// ----------------------------------------------------------------

// CreateWallet is in charge of creating a new root wallet
func CreateWallet(userId *string) *models.WalletListItem {
	bc := initBlockchain("eth")
	stamp := time.Now().UnixMilli()

	addressKeys, err := bc.GenAddrKeychain()
	if err != nil {
		exceptions.HandleAnException("<eth GenAddrKeychain> got an err: " + err.Error())
	} else {
		fmt.Println("wallet is ->\n", addressKeys)
		wt := models.EthWallet{
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

		fmt.Println("wt -> ", wt)
		// -> save wallet to main db <-
		if err := database.InsertEthWallet(&wt); err != nil {
			exceptions.HandleAnException("Database insertion got an error: " + err.Error())
		}
	}

	// returt eth address
	fmt.Println("generated eth address is\n->", addressKeys)
	return &models.WalletListItem{CoinName: "eth", Address: addressKeys.Address}
}

// GetEthereumAddressBalance -> get balance by address
func GetEthBalanceByAddress(addr string) *big.Float {
	bc := initBlockchain("eth")
	currentBalance := new(big.Float)

	// ###################################################
	// ######## DOESN"T WORK IN THE TEST-NET! ############
	// ###################################################

	addressData, err := bc.GetAddrBal(addr, nil)
	if err != nil {
		exceptions.HandleAnException("<eth GetAddrBal> got an err: " + err.Error())
	}

	currentBalance.SetString(addressData.Balance.String())
	ethValue := new(big.Float).Quo(currentBalance, big.NewFloat(math.Pow10(20)))

	// fmt.Println("balance -> ", currentBalance)
	return ethValue
}

// ===========================================================================================//
// ============================ init the blockchain connection ===============================//
// ===========================================================================================//

func initBlockchain(c string) *gobcy.API {

	bc := new(gobcy.API)
	apiToken := config.GetBitcoinAPIKey()

	if apiToken == "" {
		exceptions.HandleAnException("Init <eth> blockchain got an error. Empty API token.")
	} else {

		bc.Token = apiToken
		bc.Coin = c // options: "btc","bcy","ltc","doge","eth"
		// bc.Chain = "test3" // depend on coin: "main","test3","test"\
		bc.Chain = "main"
	}

	return bc
}
