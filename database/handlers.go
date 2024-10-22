package database

import (
	"fmt"
	"log"
	"strings"
	"wallet-cli/lib/models"
)

// doc is -> https://github.com/go-sql-driver/mysql/wiki/Examples

// InsertBtcWallet -> insert user wallet data to db
func InsertBtcWallet(dto models.BtcWallet) error {

	db := dbConnect()
	defer db.Close()

	// * pubKeys are temporary excluded  *
	sql := "INSERT INTO btcWallets (address, privateKey, publicKey, wif, scriptType, originalAddress, oapAddress, userId) VALUES (?,?,?,?,?,?,?,?)"
	res := db.QueryRow(sql, dto.Address, dto.PrivateKey, dto.PublicKey, dto.Wif, dto.ScriptType, dto.OriginalAddress, dto.OAPAddress, dto.UserId)
	fmt.Println("sql result is -> ", res)

	return res.Err()
}

// InsertEthWallet -> insert user wallet data to db
func InsertEthWallet(dto models.EthWallet) error {

	// ctx := context.Background()
	db := dbConnect()
	defer db.Close()

	// * pubKeys are temporary excluded  *
	sql := "INSERT INTO ethWallets (address, privateKey, publicKey, wif, scriptType, originalAddress, oapAddress, userId) VALUES (?,?,?,?,?,?,?,?)"
	res := db.QueryRow(sql, dto.Address, dto.PrivateKey, dto.PublicKey, dto.Wif, dto.ScriptType, dto.OriginalAddress, dto.OAPAddress, dto.UserId)
	fmt.Println("sql result is -> ", res)

	return res.Err()
}

// InsertTonWallet -> insert user wallet data to db
func InsertTonWallet(dto models.TonWallet) error {

	// ctx := context.Background()
	db := dbConnect()
	defer db.Close()

	// should be updated
	sql := "INSERT INTO tonWallets (address, addrType, privateKey, bitsLen, userId) VALUES (?,?,?,?,?)"
	res := db.QueryRow(sql, dto.Address, dto.AddrType, dto.PrivateKey, dto.BitsLen, dto.UserId)
	fmt.Println("sql result is -> ", res)

	return res.Err()
}

// InsertTrxWallet -> insert user wallet data to db
// func InsertTrxWallet(dto models.TonWallet) error {

// 	// ctx := context.Background()
// 	db := dbConnect()
// 	defer db.Close()

// 	// should be updated
// 	sql := "INSERT INTO tonWallets (address, addrType, privateKey, bitsLen, userId) VALUES (?,?,?,?,?)"
// 	res := db.QueryRow(sql, dto.Address, dto.AddrType, dto.PrivateKey, dto.BitsLen, dto.UserId)
// fmt.Println("sql result is -> ", res)

// return res.Err()
// }

// #####################################################
// ############# select area ###########################
// #####################################################

// SelectBtcPrivate -> get private key by address and userId
func SelectBtcPrivate(address string) string {

	var pKey string
	db := dbConnect()
	defer db.Close()

	sql := strings.Join([]string{"SELECT privateKey FROM btcWallets WHERE address=", "\"", address, "\""}, "")
	log.Println("sql str -> ", sql)
	err := db.QueryRow(sql).Scan(&pKey)
	if err != nil {
		panic(err.Error())
	}

	return pKey
}

// SelectTonPrivate -> get private key by address and userId
func SelectTonPrivate(address string) []byte {

	var pKey []byte
	db := dbConnect()
	defer db.Close()

	sql := strings.Join([]string{"SELECT privateKey FROM tonWallets WHERE address=", "\"", address, "\""}, "")
	err := db.QueryRow(sql).Scan(&pKey)
	if err != nil {
		panic(err.Error())
	}

	return pKey
}
