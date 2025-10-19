package theopennetwork_test

import (
	"wallet/api"
	theopennetwork "wallet/crypto-lib/the-open-network"
)

// -> if ur tests got an unexpected results - run:
// ../../database/main_test.go -> TestCleanUp to clean all data after previous testing
var (
	TEST_SVC                 = theopennetwork.InitService()
	TEST_ADDRESS             = "UQAsXIINx-qUfrzH3nJhOvm4Og-iCLIvnnL2sumivPIF4n1u" // random address from tonscan
	TEST_USER_ID       int64 = 90990                                              // also used as a peyee
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
