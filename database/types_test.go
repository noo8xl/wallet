package database

import (
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
