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
func SendSingleBtcTransaction(dto models.SendTransactionDto) string {

	var skeleton gobcy.TXSkel
	privateKey := database.SelectBtcPrivate(dto.SenderAddress)
	var err error

	amount := new(big.Int)
	amount.SetString(dto.Amount, 10)

	initBlockchain("btc")

	// use faucet to fund first
	// _, err = bc.Faucet(addressFrom, 3e5)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	//Post New TXSkeleton
	skeleton, err = bc.NewTX(gobcy.TempNewTX(dto.SenderAddress, dto.RecipientAddress, *amount), false)
	if err != nil {
		return err.Error()
	}

	// Sign it locally
	err = skeleton.Sign([]string{privateKey})
	if err != nil {
		return err.Error()
	}

	// Send TXSkeleton
	skeleton, err = bc.SendTX(skeleton)
	if err != nil {
		return err.Error()
	}
	fmt.Printf("skeleton is => %+v\n", skeleton)

	// save a tsx details to db
	return pushTransaction(skeleton.Trans.Hash)
}

// ===========================================================================================//
// ============================== function for internal usage ================================//
// ===========================================================================================//

// pushBtcTransaction -> push it to the pool in BC
func pushTransaction(hash string) string {
	// bc should be changed to BlockCypher Testnet
	// initBlockchain("btc")

	skel, err := bc.PushTX(hash)
	if err != nil {
		return err.Error()
	}
	fmt.Printf("%+v\n", skel)
	return skel.Trans.Hash
}

// // CreateAndSendTransaction creates and sends a transaction from 'from' to 'to' with the specified 'amount'.
// func CreateAndSendTransaction(from, to string, amount float64, token string) (string, error) {
// 	// Initialize the BlockCypher API client

// 	initBlockchain("btc")
// 	// Define a static fee (in BTC or relevant currency)
// 	staticFee := big.NewInt(10000) // Example static fee in satoshis (0.0001 BTC)

// 	// Convert amount to *big.Int (satoshis)
// 	// Convert from BTC to satoshis
// 	amountInSatoshis := big.NewInt(int64(amount * 1e8))

// 	// Calculate the value to send
// 	valueToSend := new(big.Int).Sub(amountInSatoshis, staticFee) // Deduct fee

// 	// Replace with the desired amount
// 	// staticFee := 0.0001 // Example static fee
// 	feeAmount := *big.NewInt(10000)

// 	// Create a new transaction
// 	tx := &gobcy.TX{
// 		Inputs: []gobcy.TXInput{
// 			{
// 				Addresses: []string{from},
// 			},
// 		},
// 		Outputs: []gobcy.TXOutput{
// 			{
// 				Addresses: []string{to},
// 				Value:     *valueToSend, // Convert to satoshis (for BTC)
// 			},
// 		},
// 		Fees: feeAmount, // Set the estimated fee in satoshis
// 	}

// 	// Send the transaction
// 	response, err := bc.NewTX(context.Background(), tx)
// 	if err != nil {
// 		return "", fmt.Errorf("failed to create transaction: %w", err)
// 	}

// 	// Return the transaction hash
// 	return response.TxHash, nil
// }
