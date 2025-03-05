package database

import (
	"testing"
)

func TestInsertBTCwalletToPermanent(t *testing.T) {

	err := TEST_SVC.InsertBtcWalletToPermanent(TEST_BTC_WT)
	if err != nil {
		t.Fatalf("permanent wallet test expected: nil; got: %v", err.Error())
	}

	err = TEST_SVC.InsertBtcWalletToOneTimeAddresses(TEST_BTC_WT)
	if err != nil {
		t.Fatalf("one time addr test expected: nil; got: %v", err.Error())
	}

}

func TestInsertETHwalletToPermanent(t *testing.T) {

	err := TEST_SVC.InsertEthWalletToPermament(TEST_ETH_WT)
	if err != nil {
		t.Fatalf("permanent wallet test expected: nil; got: %v", err.Error())
	}

	err = TEST_SVC.InsertEthWalletToOneTimeAddresses(TEST_ETH_WT)
	if err != nil {
		t.Fatalf("one time addr test expected: nil; got: %v", err.Error())
	}

}

func TestInsertTONwalletToPermanent(t *testing.T) {

	err := TEST_SVC.InsertTonWalletToPermanent(TEST_TON_WT)
	if err != nil {
		t.Fatalf("permanent wallet test expected: nil; got: %v", err.Error())
	}

	err = TEST_SVC.InsertTonWalletToOneTimeAddresses(TEST_TON_WT)
	if err != nil {
		t.Fatalf("one time addr test expected: nil; got: %v", err.Error())
	}

}

// func TestInsertTRXwalletToPermanent(t *testing.T) {

// 	err := TEST_SVC.InsertTrxWalletToPermanent(TEST_TRX_WT)
// 	if err != nil {
// 		t.Fatalf("permanent wallet test expected: nil; got: %v", err.Error())
// 	}

// 	err = TEST_SVC.InsertTrxWalletToOneTimeAddresses(TEST_TRX_WT)
// 	if err != nil {
// 		t.Fatalf("one time addr test expected: nil; got: %v", err.Error())
// 	}

// }
