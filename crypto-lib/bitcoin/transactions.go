package bitcoin

import (
	"fmt"
	"math/big"
	"wallet-cli/lib/models"

	"github.com/blockcypher/gobcy/v2"
)

// https://www.blockcypher.com/dev/?go#introduction -> doc is here <-

// CreateSingleBitcoinTransactionSkeleton -> create transaction skeleton, sign in locally and send to user for validate it
func CreateSingleBitcoinTransactionSkeleton(dto models.SendTransactionDto) gobcy.TXSkel {
	var skeleton gobcy.TXSkel
	// get addressFrom data from db
	// by userID from request <----
	var addressFrom gobcy.AddrKeychain
	var err error

	amount := new(big.Int)
	amount.SetString(dto.Amount, 0)

	initBlockchain("btc")

	// get user by dto.UserID <<<<

	//use faucet to fund first
	// _, err = bc.Faucet(addressFrom, 3e5)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	//Post New TXSkeleton
	skeleton, err = bc.NewTX(gobcy.TempNewTX(addressFrom.Address, dto.RecipientAddress, *amount), false)
	if err != nil {
		return skeleton
	}

	//Sign it locally
	err = skeleton.Sign([]string{addressFrom.Private})
	if err != nil {
		return skeleton
	}

	//Send TXSkeleton
	skeleton, err = bc.SendTX(skeleton)
	if err != nil {
		return skeleton
	}
	fmt.Printf("skeleton is => %+v\n", skeleton)

	// cache.SetTransCache(skeleton)

	return skeleton
}

// CreateMultiplyBitcoinTransactionSkeleton -> create multiply transaction skeleton
// (send coins from one address to several addresses), sign in locally and send to user for validate it
func CreateMultiplyBitcoinTransactionSkeleton(addressTo string, amount big.Int) gobcy.TXSkel {
	initBlockchain("btc")

	// >-------------- should update <--------------------

	var err error
	// get addressFrom data from db <===
	var addressFrom gobcy.AddrKeychain
	// get skeleton from cache data <<-----
	var skeleton gobcy.TXSkel

	// get user wallet address from db

	//use faucet to fund first
	// _, err = bc.Faucet(addressFrom, 3e5)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	//Post New TXSkeleton
	skeleton, err = bc.NewTX(gobcy.TempNewTX(addressFrom.Address, addressTo, amount), false)
	if err != nil {
		return skeleton
	}
	//Sign it locally
	err = skeleton.Sign([]string{addressFrom.Private})
	if err != nil {
		return skeleton
	}

	//Send TXSkeleton
	skeleton, err = bc.SendTX(skeleton)
	if err != nil {
		return skeleton
	}
	fmt.Printf("skeleton is => %+v\n", skeleton)

	// cache.SetTransCache(skeleton)

	return skeleton
}

// SendBtcTransaction -> push transaction to blockchain
func SendSingleBtcTransaction(dto models.SendTransactionDto) string {
	// create skeleton and push it to BC
	skel := CreateSingleBitcoinTransactionSkeleton(dto)
	return pushBtcTransaction(skel.Trans.Hash)
}

// ===========================================================================================//
// ============================== function for internal usage ================================//
// ===========================================================================================//

// pushBtcTransaction -> push it to the pool in BC
func pushBtcTransaction(hash string) string {
	// bc should be changed to BlockCypher Testnet
	initBlockchain("btc")

	skel, err := bc.PushTX(hash)
	if err != nil {
		return err.Error()
	}
	fmt.Printf("%+v\n", skel)
	return skel.Trans.Hash
}
