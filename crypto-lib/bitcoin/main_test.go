package bitcoin_test

import (
	"math/big"
	"testing"
	"wallet/crypto-lib/bitcoin"
)

func TestInitService(t *testing.T) {
	svc := bitcoin.InitService()
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
	balance, err := TEST_SVC.GetBalanceByAddress(TEST_ADDRESS)
	if err != nil {
		t.Fatalf("TestGetBalanceData expected non-empty float; got an error: %v", err)
	}
	result := balance.Cmp(zero)
	if result == -1 {
		t.Fatalf("TestGetBalanceData expected positive non-zero value; got: %v", balance)
	}
	if result == 1 {
		t.Fatalf("TestGetBalanceData expected  non-zero value; got: %v", balance)
	}

	t.Logf("TestGetBalanceData result -> %v", balance)
}
