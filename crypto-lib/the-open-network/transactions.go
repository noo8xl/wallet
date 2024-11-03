package theopennetwork

import (
	"fmt"
	"wallet-cli/lib/models"

	"github.com/xssnick/tonutils-go/ton"
)

func SendSingleTonTransaction(dto *models.SendTransactionDto) string {
	// manage smth and send trx ->
	tonAPI := initTonAPIConnection()

	return signAndPushTransaction(dto, tonAPI)
}

// ===========================================================================================//
// ============================== function for internal usage ================================//
// ===========================================================================================//

func signAndPushTransaction(dto *models.SendTransactionDto, api *ton.APIClient) string {
	fmt.Println("dto ->\n", dto)
	fmt.Println("api -> ", api)
	// privateKey := database.SelectTonPrivate(dto.SenderAddress)

	// Return the transaction hash
	return "tsx hash"
}
