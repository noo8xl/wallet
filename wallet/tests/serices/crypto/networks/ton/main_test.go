package theopennetwork_test

import (
	"math/big"
	"testing"
	theopennetwork "wallet/crypto-lib/the-open-network"
)

func TestInitService(t *testing.T) {
	svc := theopennetwork.InitService()
	if svc == nil {
		t.Fatalf("cannot init service. Got %v instead of &theopennetwork.Service{}", svc)
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
	// https://tonscan.org/address/UQAsXIINx-qUfrzH3nJhOvm4Og-iCLIvnnL2sumivPIF4n1u

	zero := big.NewFloat(0)
	balance, err := TEST_SVC.GetBalanceByAddress(TEST_ADDRESS)
	result := balance.Cmp(zero)
	if result == -1 || err != nil {
		t.Fatalf("TestGetBalanceData expected positive value != 0; got: %v", err)
	}
	t.Logf("TestGetBalanceData result -> %v", balance)
}
