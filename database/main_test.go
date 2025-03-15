package database

import (
	"testing"
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

	err = TEST_SVC.deleteTestWalletItem("trxWallet")
	if err != nil {
		t.Fatalf("TestCleanUp returns an error: %v", err)
	}
	err = TEST_SVC.deleteTestWalletItem("oneTimeTrxWallets")
	if err != nil {
		t.Fatalf("TestCleanUp returns an error: %v", err)
	}

	err = TEST_SVC.deleteTestWalletItem("solWallet")
	if err != nil {
		t.Fatalf("TestCleanUp returns an error: %v", err)
	}
	err = TEST_SVC.deleteTestWalletItem("oneTimeSolWallets")
	if err != nil {
		t.Fatalf("TestCleanUp returns an error: %v", err)
	}

	err = TEST_SVC.deleteTestWalletItem("dogeWallet")
	if err != nil {
		t.Fatalf("TestCleanUp returns an error: %v", err)
	}
	err = TEST_SVC.deleteTestWalletItem("oneTimeDogeWallets")
	if err != nil {
		t.Fatalf("TestCleanUp returns an error: %v", err)
	}

	err = TEST_SVC.deleteTestWalletItem("ltcWallet")
	if err != nil {
		t.Fatalf("TestCleanUp returns an error: %v", err)
	}
	err = TEST_SVC.deleteTestWalletItem("oneTimeLtcWallets")
	if err != nil {
		t.Fatalf("TestCleanUp returns an error: %v", err)
	}

}
