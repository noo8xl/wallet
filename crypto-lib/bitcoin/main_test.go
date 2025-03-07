package bitcoin_test

import (
	"math/big"
	"testing"
	"wallet/api"
	"wallet/crypto-lib/bitcoin"
)

// -> if ur tests got an unexpected results - run:
// ../../database/main_test.go -> TestCleanUp to clean all data after previous testing
var (
	TEST_SVC                 = bitcoin.InitBitcoinService()
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

func TestInitService(t *testing.T) {
	svc := bitcoin.InitBitcoinService()
	if svc == nil {
		t.Fatalf("cannot init service. Got %v instead of &bitcoin.Service{}", svc)
	}
}

func TestCreatePermanentWallet(t *testing.T) {

	wt := TEST_SVC.CreatePermanentWallet(TEST_USER_ID)
	if wt == nil {
		t.Fatalf("TestCreatePermanentWallet 1 expected &api.WalletItem, got: %v", wt)
	}

	wt = TEST_SVC.CreatePermanentWallet(TEST_USER_ID)
	if wt != nil {
		t.Fatalf("TestCreatePermanentWallet 2 expected nil, got: %v", wt)
	}

	wt2 := TEST_SVC.CreatePermanentWallet(TEST_BENEFICIAR_ID)
	if wt2 == nil {
		t.Fatalf("TestCreatePermanentWallet 1 expected &api.WalletItem, got: %v", wt2)
	}

	TEST_PEYEE_ADDRESS = wt.Address
	TEST_BENEFICIAR_ADDRESS = wt.Address

}

func TestCreateOneTimeAddress(t *testing.T) {

	nilItem := TEST_SVC.CreateOneTimeddress(TEST_USER_ID)
	if nilItem == nil {
		t.Fatalf("TestCreateOneTimeAddress 1 expected &api.WalletItem, got: %v", nilItem)
	}

	nilItem = TEST_SVC.CreateOneTimeddress(TEST_USER_ID)
	if nilItem == nil {
		t.Fatalf("TestCreateOneTimeAddress 2 expected nil, got: %v", nilItem)
	}

}

func TestGetBalanceData(t *testing.T) {
	// -> to run this test u should change bc networt to the "main" in main.go file in this package *
	// -> to verify the value (if nessesary)
	// https://www.blockchain.com/explorer/addresses/btc/bc1qxy2kgdygjrsqtzq2n0yrf2493p83kkfjhx0wlh

	zero := big.NewFloat(0)
	balance := TEST_SVC.GetBalanceByAddress(TEST_ADDRESS)
	result := balance.Cmp(zero)
	if result == -1 {
		t.Fatalf("TestGetBalanceData expected positive value != 0; got: %v", balance)
	}
	t.Logf("TestGetBalanceData result -> %v", balance)
}

func TestSingleTransaction(t *testing.T) {
	TEST_SINGLE_TSX_DTO.Payee.PeyeeAddress = TEST_PEYEE_ADDRESS
	TEST_SINGLE_TSX_DTO.Beneficiar.BeneficiarAddress = TEST_BENEFICIAR_ADDRESS

	hash := TEST_SVC.SendSingleTransaction(TEST_SINGLE_TSX_DTO)
	if hash == "" {
		t.Fatalf("TestSingleTransaction expected non-empty string; got: %v", hash)
	}

	t.Logf("TestSingleTransaction transaction hash -> %v", hash)
}

// func TestMultTransaction(t *testing.T) {

// }
