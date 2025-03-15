package bitcoin_test

import (
	"wallet/api"
	"wallet/crypto-lib/bitcoin"
)

// -> if ur tests got an unexpected results - run:
// ../../database/main_test.go -> TestCleanUp to clean all data after previous testing
var (
	TEST_SVC                 = bitcoin.InitService()
	TEST_ADDRESS             = "bc1qxy2kgdygjrsqtzq2n0yrf2493p83kkfjhx0wlh" // random address from the blockchain explorer
	TEST_USER_ID       int64 = 90990                                        // also used as a peyee
	TEST_BENEFICIAR_ID int64 = 8831721

	TEST_BENEFICIAR_ADDRESS = ""
	TEST_PEYEE_ADDRESS      = ""

	TEST_SINGLE_TSX_DTO = &api.SendSingleTsxRequest{
		CoinName: "btc",
		Payee: &api.PeyeeData{
			PeyeeAddress: "",
			PeyeeId:      TEST_USER_ID,
			Amount:       "0.023",
		},
		Beneficiar: &api.BeneficiarData{
			BeneficiarAddress: "",
			BeneficiarId:      TEST_BENEFICIAR_ID,
			Amount:            "0.023",
		},
	}
)
