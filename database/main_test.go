package database

import (
	"testing"
	"time"
	"wallet/lib/models"
)

var (
	TEST_SVC = InitDbService()

	TEST_BTC_WT = &models.BtcWallet{
		Address:         "test btc address here ",
		PrivateKey:      "Private key here",
		PublicKey:       "and public key also",
		Wif:             "now -> WIF",
		ScriptType:      "scrypt type val",
		OriginalAddress: "some str here",
		OAPAddress:      "unknown field ?",
		CreatedAt:       time.Now().UnixMilli(),
		UpdatedAt:       time.Now().UnixMilli(),
		CustomerId:      90990,
	}

	TEST_ETH_WT = &models.EthWallet{
		Address:         "test btc address here ",
		PrivateKey:      "Private key here",
		PublicKey:       "and public key also",
		Wif:             "now -> WIF",
		ScriptType:      "scrypt type val",
		OriginalAddress: "some str here",
		OAPAddress:      "unknown field ?",
		CreatedAt:       time.Now().UnixMilli(),
		UpdatedAt:       time.Now().UnixMilli(),
		CustomerId:      90990,
	}

	TEST_TRX_WT = &models.TrxWallet{
		Address:    "test btc address here ",
		PrivateKey: "Private key here",
		PublicKey:  "and public key also",
		Wif:        "now -> WIF",
		CreatedAt:  time.Now().UnixMilli(),
		UpdatedAt:  time.Now().UnixMilli(),
		CustomerId: 90990,
	}

	TEST_TON_WT = &models.TonWallet{
		Address:    "test btc address here ",
		AddrType:   123,
		PrivateKey: []byte("Some private value"),
		BitsLen:    932,
		CreatedAt:  time.Now().UnixMilli(),
		UpdatedAt:  time.Now().UnixMilli(),
		CustomerId: 90990,
	}

	TEST_TON_ONE_TIME_ADDRESS = &models.TonOneTimeWallet{
		Address:    "test btc address here ",
		AddrType:   123,
		PrivateKey: []byte("Some private value"),
		BitsLen:    932,
	}
)

func TestConnection(t *testing.T) {
	conn := InitDbService()
	conn.db = dbConnect()
	if conn.db == nil {
		t.Fatalf("expected: *DatabaseService NOT NULL struct; got: %v", conn.db)
	}
	defer conn.db.Close()
}

func TestCleanUp(t *testing.T) {

	err := TEST_SVC.deleteTestWalletItem("btcWallets")
	if err != nil {
		t.Fatalf("TestCleanUp returns an error: %v", err)
	}
	err = TEST_SVC.deleteTestWalletItem("oneTimeBtcWallets")
	if err != nil {
		t.Fatalf("TestCleanUp returns an error: %v", err)
	}

	err = TEST_SVC.deleteTestWalletItem("ethWallets")
	if err != nil {
		t.Fatalf("TestCleanUp returns an error: %v", err)
	}
	err = TEST_SVC.deleteTestWalletItem("oneTimeEthWallets")
	if err != nil {
		t.Fatalf("TestCleanUp returns an error: %v", err)
	}

	err = TEST_SVC.deleteTestWalletItem("tonWallets")
	if err != nil {
		t.Fatalf("TestCleanUp returns an error: %v", err)
	}
	err = TEST_SVC.deleteTestWalletItem("oneTimeTonWallets")
	if err != nil {
		t.Fatalf("TestCleanUp returns an error: %v", err)
	}

	// TEST_SVC.deleteTestWalletItem("trxWallet")
	// TEST_SVC.deleteTestWalletItem("oneTimeTrxWallets")

	// TEST_SVC.deleteTestWalletItem("solWallet")
	// TEST_SVC.deleteTestWalletItem("oneTimeSolWallets")

}
