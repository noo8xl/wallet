package database

import "testing"

func TestSelectBtcPrivate(t *testing.T) {
	pk := TEST_SVC.SelectBtcPrivate(TEST_BTC_WT.Address)
	if pk == "" {
		t.Fatalf("expected value: string NOT NULL val; got: %s", pk)
	}
	// t.Logf("NOT NULL string public key is: %s", pk)
}

func TestSelectEthPrivate(t *testing.T) {
	pk := TEST_SVC.SelectEthPrivate(TEST_ETH_WT.Address)
	if pk == "" {
		t.Fatalf("expected value: string NOT NULL val; got: %s", pk)
	}
	// t.Logf("NOT NULL string public key is: %v", pk)
}

func TestSelectTonPrivate(t *testing.T) {
	pk := TEST_SVC.SelectTonPrivate(TEST_TON_WT.Address)
	if string(pk) == "" {
		t.Fatalf("expected keyvalue: []byte NOT NULL val; got: %s", pk)
	}
	// t.Logf("NOT NULL []byte public key is: %s", pk)
}

func TestDefineAndGetTableNameByBlockchainShortName(t *testing.T) {

	tableName := TEST_SVC.defineAndGetTableNameByBlockchainShortName("btc")
	if tableName != "btcWallets" {
		t.Fatalf("btc blockchain expected: <btcWallets> string val; got: %s", tableName)
	}

	tableName = TEST_SVC.defineAndGetTableNameByBlockchainShortName("eth")
	if tableName != "ethWallets" {
		t.Fatalf("eth blockchain expected: <ethWallets> string val; got: %s", tableName)
	}

	tableName = TEST_SVC.defineAndGetTableNameByBlockchainShortName("ton")
	if tableName != "tonWallets" {
		t.Fatalf("ton blockchain expected: <tonWallets> string val; got: %s", tableName)
	}

	// tableName = TEST_SVC.defineAndGetTableNameByBlockchainShortName("trx")
	// if tableName != "trxWallets" {
	// 	t.Fatalf("trx blockchain expected: <trxWallets> string val; got: %s", tableName)
	// }

	// tableName = TEST_SVC.defineAndGetTableNameByBlockchainShortName("sol")
	// if tableName != "solWallets" {
	// 	t.Fatalf("sol blockchain expected: <solWallets> string val; got: %s", tableName)
	// }

}

func TestIWalletExists(t *testing.T) {

	isExists := TEST_SVC.IsWalletExists(90990, "btc")
	if !isExists {
		t.Fatalf("btc blockchain expected: bool true val; got: %v", isExists)
	}

	isExists = TEST_SVC.IsWalletExists(90990, "eth")
	if !isExists {
		t.Fatalf("eth blockchain expected: bool true val; got: %v", isExists)
	}

	isExists = TEST_SVC.IsWalletExists(90990, "ton")
	if !isExists {
		t.Fatalf("ton blockchain expected: bool true val; got: %v", isExists)
	}

	// isExists = TEST_SVC.IsWalletExists(90990, "trx")
	// if !isExists {
	// 	t.Fatalf("trx blockchain expected: bool true val; got: %v", isExists)
	// }

	// isExists = TEST_SVC.IsWalletExists(90990, "sol")
	// if !isExists {
	// 	t.Fatalf("sol blockchain expected: bool true val; got: %v", isExists)
	// }
}
