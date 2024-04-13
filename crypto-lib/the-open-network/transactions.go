package theopennetwork

import (
	"fmt"
	"wallet-cli/lib/models"
)

func SendSingleTonTransaction(dto models.SendTransactionDto) string {
	// manage smth and send trx ->
	return sendTransaction(dto)
}

// ===========================================================================================//
// ============================== function for internal usage ================================//
// ===========================================================================================//

func sendTransaction(dto models.SendTransactionDto) string {
	fmt.Println("dto ->\n", dto)
	return "test str hash"
}
