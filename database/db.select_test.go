package database

import "testing"

func TestSelectBtcPrivate(t *testing.T) {
	var err error

	_, err = TEST_SVC.DefineaBlockchainAndGetPrivateKey(TEST_BTC_WT.Address, "btc", 0)
	if err != nil {
		t.Fatalf("expected value: string NOT NULL val; got an error: %v", err)
	}

	_, err = TEST_SVC.DefineaBlockchainAndGetPrivateKey(TEST_BTC_WT.Address, "btc", 1)
	if err != nil {
		t.Fatalf("expected value: string NOT NULL val; got an error: %v", err)
	}

}

func TestSelectEthPrivate(t *testing.T) {
	var err error

	_, err = TEST_SVC.DefineaBlockchainAndGetPrivateKey(TEST_BTC_WT.Address, "eth", 0)
	if err != nil {
		t.Fatalf("expected value: string NOT NULL val; got an error: %v", err)
	}

	_, err = TEST_SVC.DefineaBlockchainAndGetPrivateKey(TEST_BTC_WT.Address, "eth", 1)
	if err != nil {
		t.Fatalf("expected value: string NOT NULL val; got an error: %v", err)
	}

}

func TestSelectTonPrivate(t *testing.T) {
	var err error

	_, err = TEST_SVC.DefineaBlockchainAndGetPrivateKey(TEST_BTC_WT.Address, "ton", 0)
	if err != nil {
		t.Fatalf("expected value: string NOT NULL val; got an error: %v", err)
	}

	_, err = TEST_SVC.DefineaBlockchainAndGetPrivateKey(TEST_BTC_WT.Address, "ton", 1)
	if err != nil {
		t.Fatalf("expected value: string NOT NULL val; got an error: %v", err)
	}

}

func TestSelectTrxPrivate(t *testing.T) {
	var err error
	_, err = TEST_SVC.DefineaBlockchainAndGetPrivateKey(TEST_BTC_WT.Address, "trx", 0)
	if err != nil {
		t.Fatalf("expected value: string NOT NULL val; got an error: %v", err)
	}

	_, err = TEST_SVC.DefineaBlockchainAndGetPrivateKey(TEST_BTC_WT.Address, "trx", 1)
	if err != nil {
		t.Fatalf("expected value: string NOT NULL val; got an error: %v", err)
	}

}

func TestSelectSolPrivate(t *testing.T) {
	var err error

	_, err = TEST_SVC.DefineaBlockchainAndGetPrivateKey(TEST_BTC_WT.Address, "sol", 0)
	if err != nil {
		t.Fatalf("expected value: string NOT NULL val; got an error: %v", err)
	}

	_, err = TEST_SVC.DefineaBlockchainAndGetPrivateKey(TEST_BTC_WT.Address, "sol", 1)
	if err != nil {
		t.Fatalf("expected value: string NOT NULL val; got an error: %v", err)
	}

}

func TestSelectDogePrivate(t *testing.T) {
	var err error

	_, err = TEST_SVC.DefineaBlockchainAndGetPrivateKey(TEST_BTC_WT.Address, "doge", 0)
	if err != nil {
		t.Fatalf("expected value: string NOT NULL val; got an error: %v", err)
	}

	_, err = TEST_SVC.DefineaBlockchainAndGetPrivateKey(TEST_BTC_WT.Address, "doge", 1)
	if err != nil {
		t.Fatalf("expected value: string NOT NULL val; got an error: %v", err)
	}

}

func TestSelectLtcPrivate(t *testing.T) {
	var err error

	_, err = TEST_SVC.DefineaBlockchainAndGetPrivateKey(TEST_BTC_WT.Address, "ltc", 0)
	if err != nil {
		t.Fatalf("expected value: string NOT NULL val; got an error: %v", err)
	}

	_, err = TEST_SVC.DefineaBlockchainAndGetPrivateKey(TEST_BTC_WT.Address, "ltc", 1)
	if err != nil {
		t.Fatalf("expected value: string NOT NULL val; got an error: %v", err)
	}
}

func TestSelectaWrongCoinNameAndOpts(t *testing.T) {
	var err error
	var val string

	val, err = TEST_SVC.DefineaBlockchainAndGetPrivateKey(TEST_BTC_WT.Address, "rom", 0)
	if err == nil {
		t.Fatalf("expected value: coin name error; got: %v", val)
	}

	val, err = TEST_SVC.DefineaBlockchainAndGetPrivateKey(TEST_BTC_WT.Address, "trx", 9)
	if val != "" || err == nil {
		t.Fatalf("expected value: invalid opts error; got: %v", val)
	}
}

func TestDefineAndGetTableNameByBlockchainShortName(t *testing.T) {
	var err error

	_, err = TEST_SVC.defineAndGetPermanentTableNameByBlockchainShortName("btc")
	if err != nil {
		t.Fatalf("btc blockchain expected: <btcWallets> string val; got: %v", err)
	}

	_, err = TEST_SVC.defineAndGetOneTimeTableNameByBlockchainShortName("btc")
	if err != nil {
		t.Fatalf("btc blockchain expected: <oneTimeBtcWallets> string val; got: %v", err)
	}

	_, err = TEST_SVC.defineAndGetPermanentTableNameByBlockchainShortName("eth")
	if err != nil {
		t.Fatalf("eth blockchain expected: <ethWallets> string val; got: %v", err)
	}

	_, err = TEST_SVC.defineAndGetOneTimeTableNameByBlockchainShortName("eth")
	if err != nil {
		t.Fatalf("eth blockchain expected: <oneTimeEthWallets> string val; got: %v", err)
	}

	_, err = TEST_SVC.defineAndGetPermanentTableNameByBlockchainShortName("ltc")
	if err != nil {
		t.Fatalf("ltc blockchain expected: <ltcWallets> string val; got: %v", err)
	}

	_, err = TEST_SVC.defineAndGetOneTimeTableNameByBlockchainShortName("ltc")
	if err != nil {
		t.Fatalf("ltc blockchain expected: <oneTimeLtcWallets> string val; got: %v", err)
	}

	_, err = TEST_SVC.defineAndGetPermanentTableNameByBlockchainShortName("doge")
	if err != nil {
		t.Fatalf("doge blockchain expected: <dogeWallets> string val; got: %v", err)
	}

	_, err = TEST_SVC.defineAndGetOneTimeTableNameByBlockchainShortName("doge")
	if err != nil {
		t.Fatalf("doge blockchain expected: <oneTimeDogeWallets> string val; got: %v", err)
	}

	_, err = TEST_SVC.defineAndGetPermanentTableNameByBlockchainShortName("ton")
	if err != nil {
		t.Fatalf("ton blockchain expected: <tonWallets> string val; got: %v", err)
	}

	_, err = TEST_SVC.defineAndGetOneTimeTableNameByBlockchainShortName("ton")
	if err != nil {
		t.Fatalf("ton blockchain expected: <tonWallets> string val; got: %v", err)
	}

	_, err = TEST_SVC.defineAndGetPermanentTableNameByBlockchainShortName("trx")
	if err != nil {
		t.Fatalf("trx blockchain expected: <trxWallets> string val; got: %v", err)
	}

	_, err = TEST_SVC.defineAndGetOneTimeTableNameByBlockchainShortName("trx")
	if err != nil {
		t.Fatalf("trx blockchain expected: <oneTimeTrxWallets> string val; got: %v", err)
	}

	_, err = TEST_SVC.defineAndGetPermanentTableNameByBlockchainShortName("sol")
	if err != nil {
		t.Fatalf("sol blockchain expected: <solWallets> string val; got: %v", err)
	}

	_, err = TEST_SVC.defineAndGetOneTimeTableNameByBlockchainShortName("sol")
	if err != nil {
		t.Fatalf("sol blockchain expected: <oneTimaSolWallets> string val; got: %v", err)
	}

}

func TestIWalletExists(t *testing.T) {

	var err error

	_, err = TEST_SVC.IsWalletExists(90990, "btc", 0)
	if err != nil {
		t.Fatalf("btc blockchain expected: bool true val; got an error: %v", err)
	}

	_, err = TEST_SVC.IsWalletExists(90990, "btc", 1)
	if err != nil {
		t.Fatalf("btc blockchain expected: bool true val; got an error: %v", err)
	}

	_, err = TEST_SVC.IsWalletExists(90990, "eth", 0)
	if err != nil {
		t.Fatalf("eth blockchain expected: bool true val; got: %v", err)
	}

	_, err = TEST_SVC.IsWalletExists(90990, "eth", 1)
	if err != nil {
		t.Fatalf("eth blockchain expected: bool true val; got: %v", err)
	}

	_, err = TEST_SVC.IsWalletExists(90990, "doge", 0)
	if err != nil {
		t.Fatalf("doge blockchain expected: bool true val; got: %v", err)
	}

	_, err = TEST_SVC.IsWalletExists(90990, "doge", 1)
	if err != nil {
		t.Fatalf("doge blockchain expected: bool true val; got: %v", err)
	}

	_, err = TEST_SVC.IsWalletExists(90990, "ltc", 0)
	if err != nil {
		t.Fatalf("ltc blockchain expected: bool true val; got: %v", err)
	}

	_, err = TEST_SVC.IsWalletExists(90990, "ltc", 1)
	if err != nil {
		t.Fatalf("ltc blockchain expected: bool true val; got: %v", err)
	}

	_, err = TEST_SVC.IsWalletExists(90990, "ton", 0)
	if err != nil {
		t.Fatalf("ton blockchain expected: bool true val; got: %v", err)
	}

	_, err = TEST_SVC.IsWalletExists(90990, "ton", 1)
	if err != nil {
		t.Fatalf("ton blockchain expected: bool true val; got: %v", err)
	}

	_, err = TEST_SVC.IsWalletExists(90990, "trx", 0)
	if err != nil {
		t.Fatalf("trx blockchain expected: bool true val; got: %v", err)
	}

	_, err = TEST_SVC.IsWalletExists(90990, "trx", 1)
	if err != nil {
		t.Fatalf("trx blockchain expected: bool true val; got: %v", err)
	}

	_, err = TEST_SVC.IsWalletExists(90990, "sol", 0)
	if err != nil {
		t.Fatalf("sol blockchain expected: bool true val; got: %v", err)
	}

	_, err = TEST_SVC.IsWalletExists(90990, "sol", 1)
	if err != nil {
		t.Fatalf("sol blockchain expected: bool true val; got: %v", err)
	}

}
