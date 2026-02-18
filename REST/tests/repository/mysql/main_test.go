package mysql

import (
	"testing"
)

func TestCleanUp(t *testing.T) {

	err := TEST_SVC.DeleteTestWalletItem("btcWallets")
	if err != nil {
		t.Fatalf("TestCleanUp returns an error: %v", err)
	}
	err = TEST_SVC.DeleteTestWalletItem("oneTimeBtcWallets")
	if err != nil {
		t.Fatalf("TestCleanUp returns an error: %v", err)
	}

	err = TEST_SVC.DeleteTestWalletItem("ethWallets")
	if err != nil {
		t.Fatalf("TestCleanUp returns an error: %v", err)
	}
	err = TEST_SVC.DeleteTestWalletItem("oneTimeEthWallets")
	if err != nil {
		t.Fatalf("TestCleanUp returns an error: %v", err)
	}

	err = TEST_SVC.DeleteTestWalletItem("tonWallets")
	if err != nil {
		t.Fatalf("TestCleanUp returns an error: %v", err)
	}
	err = TEST_SVC.DeleteTestWalletItem("oneTimeTonWallets")
	if err != nil {
		t.Fatalf("TestCleanUp returns an error: %v", err)
	}

	err = TEST_SVC.DeleteTestWalletItem("trxWallet")
	if err != nil {
		t.Fatalf("TestCleanUp returns an error: %v", err)
	}
	err = TEST_SVC.DeleteTestWalletItem("oneTimeTrxWallets")
	if err != nil {
		t.Fatalf("TestCleanUp returns an error: %v", err)
	}

	err = TEST_SVC.DeleteTestWalletItem("solWallet")
	if err != nil {
		t.Fatalf("TestCleanUp returns an error: %v", err)
	}
	err = TEST_SVC.DeleteTestWalletItem("oneTimeSolWallets")
	if err != nil {
		t.Fatalf("TestCleanUp returns an error: %v", err)
	}

	err = TEST_SVC.DeleteTestWalletItem("dogeWallet")
	if err != nil {
		t.Fatalf("TestCleanUp returns an error: %v", err)
	}
	err = TEST_SVC.DeleteTestWalletItem("oneTimeDogeWallets")
	if err != nil {
		t.Fatalf("TestCleanUp returns an error: %v", err)
	}

	err = TEST_SVC.DeleteTestWalletItem("ltcWallet")
	if err != nil {
		t.Fatalf("TestCleanUp returns an error: %v", err)
	}
	err = TEST_SVC.DeleteTestWalletItem("oneTimeLtcWallets")
	if err != nil {
		t.Fatalf("TestCleanUp returns an error: %v", err)
	}

}
