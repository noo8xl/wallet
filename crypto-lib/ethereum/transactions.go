package ethereum

import (
	"fmt"
	"math/big"
	"wallet-cli/database"
	"wallet-cli/lib/exceptions"
	"wallet-cli/lib/models"

	"github.com/blockcypher/gobcy/v2"
)

func SendSingleEthTransaction(dto *models.SendTransactionDto) string {

	var skeleton gobcy.TXSkel
	privateKey := database.SelectBtcPrivate(dto.SenderAddress)
	var err error

	amount := new(big.Int)
	amount.SetString(dto.Amount, 10)

	bc := initBlockchain("btc")

	// use faucet to fund first
	// _, err = bc.Faucet(addressFrom, 3e5)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	//Post New TXSkeleton
	skeleton, err = bc.NewTX(gobcy.TempNewTX(dto.SenderAddress, dto.RecipientAddress, *amount), false)
	if err != nil {
		exceptions.HandleAnException("<eth create transactions> got an err: " + err.Error())
	}

	// Sign it locally
	err = skeleton.Sign([]string{privateKey})
	if err != nil {
		exceptions.HandleAnException("<eth sign transactions> got an err: " + err.Error())
	}

	// Send TXSkeleton
	skeleton, err = bc.SendTX(skeleton)
	if err != nil {
		exceptions.HandleAnException("<eth send transactions> got an err: " + err.Error())
	}
	fmt.Printf("skeleton is => %+v\n", skeleton)

	// save a tsx details to db
	return pushTransaction(skeleton.Trans.Hash, bc)
}

// ===========================================================================================//
// ============================== function for internal usage ================================//
// ===========================================================================================//

// pushBtcTransaction -> push it to the pool in BC
func pushTransaction(hash string, bc *gobcy.API) string {
	// bc should be changed to BlockCypher Testnet
	// initBlockchain("btc")

	skel, err := bc.PushTX(hash)
	if err != nil {
		exceptions.HandleAnException("<eth push transactions> got an err: " + err.Error())
	}
	fmt.Printf("%+v\n", skel)
	return skel.Trans.Hash
}
