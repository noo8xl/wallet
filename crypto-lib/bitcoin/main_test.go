package bitcoin_test

import (
	"testing"
	"wallet/crypto-lib/bitcoin"
)

func TestInitService(t *testing.T) {
	btcService := bitcoin.InitBitcoinService()
	if btcService == nil {
		t.Errorf("cannot init btc service. got %v instead", btcService)
	}
}

// func TestGenerateAddress(t *testing.T) {
// 	btcService := bitcoin.InitBitcoinService()
// 	item := btcService.CreatePermanentWallet(1)
// 	if item == nil {
// 		t.Errorf("got an arror at CreateWallet")
// 	}

// 	log.Printf("got an item: %v", item)
// }

// func TestCreatePermanentWallet(t *testing.T) {

// }

// func TestSingleTransaction(t *testing.T) {

// }

// func TestMultTransaction(t *testing.T) {

// }
