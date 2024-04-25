package bitcoin

import (
	"fmt"
	"math/big"
	"wallet-cli/database"
	"wallet-cli/lib/models"

	"github.com/blockcypher/gobcy/v2"
)

// https://www.blockcypher.com/dev/?go#introduction -> doc is here <-

// CreateSingleBitcoinTransactionSkeleton -> create transaction skeleton, sign in locally and send to user for validate it
func CreateSingleBitcoinTransactionSkeleton(dto models.SendTransactionDto) string {

	//  -> https://medium.com/coinmonks/estep-by-step-create-and-broadcast-a-bitcoin-transaction-on-testnet-588daacc2b7a

	var skeleton gobcy.TXSkel
	// get addressFrom data from db
	// by userID from request <----
	privateKey := database.SelectBtcPrivate(dto.SenderAddress)
	var err error

	amount := new(big.Int)
	amount.SetString(dto.Amount, 0)

	initBlockchain("btc")

	//use faucet to fund first
	// _, err = bc.Faucet(addressFrom, 3e5)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	//Post New TXSkeleton
	skeleton, err = bc.NewTX(gobcy.TempNewTX(dto.SenderAddress, dto.RecipientAddress, *amount), false)
	if err != nil {
		return err.Error()
	}

	//Sign it locally
	err = skeleton.Sign([]string{privateKey})
	if err != nil {
		return err.Error()
	}

	//Send TXSkeleton
	skeleton, err = bc.SendTX(skeleton)
	if err != nil {
		return err.Error()
	}
	fmt.Printf("skeleton is => %+v\n", skeleton)

	// cache.SetTransCache(skeleton)
	return ""
	// return pushBtcTransaction(skeleton.Trans.Hash)
}

// CreateMultiplyBitcoinTransactionSkeleton -> create multiply transaction skeleton
// (send coins from one address to several addresses), sign in locally and send to user for validate it
func CreateMultiplyBitcoinTransactionSkeleton(dto models.SendTransactionDto) string {
	initBlockchain("btc")

	// >-------------- should update <--------------------
	// https://medium.com/coinmonks/estep-by-step-create-and-broadcast-a-bitcoin-transaction-on-testnet-588daacc2b7a

	var err error
	privateKey := database.SelectBtcPrivate(dto.SenderAddress)
	// get skeleton from cache data <<-----
	var skeleton gobcy.TXSkel
	amount := new(big.Int)
	amount.SetString(dto.Amount, 0)

	//use faucet to fund first
	// _, err = bc.Faucet(addressFrom, 3e5)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	//Post New TXSkeleton
	skeleton, err = bc.NewTX(gobcy.TempNewTX(dto.SenderAddress, dto.RecipientAddress, *amount), false)
	if err != nil {
		return err.Error()
	}
	//Sign it locally
	err = skeleton.Sign([]string{privateKey})
	if err != nil {
		return err.Error()
	}

	//Send TXSkeleton
	skeleton, err = bc.SendTX(skeleton)
	if err != nil {
		return err.Error()
	}
	fmt.Printf("skeleton is => %+v\n", skeleton)

	// cache.SetTransCache(skeleton)
	// pushBtcTransaction(skel.Trans.Hash)
	return skeleton.Trans.Hash
}

// // SendBtcTransaction -> push transaction to blockchain
// func SendSingleBtcTransaction(dto models.SendTransactionDto) string {
// 	// create skeleton and push it to BC
// 	skel := CreateSingleBitcoinTransactionSkeleton(dto)
// 	return pushBtcTransaction(skel.Trans.Hash)
// }

// ===========================================================================================//
// ============================== function for internal usage ================================//
// ===========================================================================================//

// // pushBtcTransaction -> push it to the pool in BC
// func pushBtcTransaction(hash string) string {
// 	// bc should be changed to BlockCypher Testnet
// 	initBlockchain("btc")

// 	skel, err := bc.PushTX(hash)
// 	if err != nil {
// 		return err.Error()
// 	}
// 	fmt.Printf("%+v\n", skel)
// 	return skel.Trans.Hash
// }
