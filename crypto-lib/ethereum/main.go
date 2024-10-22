// Package ethereumwallet -> is eth network interacting
package ethereum

import (
	"fmt"
	"log"
	"math"
	"math/big"
	"os"
	"time"
	"wallet-cli/config"
	"wallet-cli/database"
	"wallet-cli/lib/models"

	"github.com/blockcypher/gobcy/v2"
)

var bc gobcy.API
var apiToken string = config.GetEthereumAPIKey()

// ----------------------------------------------------------------

// CreateWallet is in charge of creating a new root wallet
func CreateWallet(userID string) string {
	initBlockchain("eth")
	stamp := time.Now().UnixMilli()

	addressKeys, err := bc.GenAddrKeychain()
	if err != nil {
		panic(err)
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
			UserId:          userID,
			// PubKeys:         addressKeys.PubKeys,
		}

		fmt.Println("wt -> ", wt)
		// -> save wallet to main db <-
		if err := database.InsertEthWallet(wt); err != nil {
			log.Println("database insertion error", err)
			os.Exit(1)
		}
	}

	// returt eth address
	fmt.Println("generated eth address is\n->", addressKeys)
	return addressKeys.Address
}

// GetEthereumAddressBalance -> get balance by address
func GetEthBalanceByAddress(addr string) *big.Float {
	initBlockchain("eth")
	currentBalance := new(big.Float)

	// ###################################################
	// ######## DOESN"T WORK IN THE TEST-NET! ############
	// ###################################################

	addressData, err := bc.GetAddrBal(addr, nil)
	if err != nil {
		fmt.Println("err is -> ", err)
	}

	currentBalance.SetString(addressData.Balance.String())
	ethValue := new(big.Float).Quo(currentBalance, big.NewFloat(math.Pow10(20)))

	// fmt.Println("balance -> ", currentBalance)
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
		bc.Chain = "main" //depend on coin: "main","test3","test"

		// fmt.Println("blockchain  ->", bc)
	}

}
