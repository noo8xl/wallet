package database

import (
	"fmt"
	"time"
	"wallet-cli/lib/models"
)

// CreateUserWallet -> insert user wallet data to db
func CreateUserBtcWallet(email string, walletDto models.BtcWallet) {

	db := dbConnect()
	defer db.Close()
	stamp := time.Now().UnixMilli()

	wt := models.BtcWallet{
		PublicKey:  walletDto.PublicKey,
		PrivateKey: walletDto.PrivateKey,
		CreatedAt:  stamp,
		UpdatedAt:  stamp,
		UserID:     walletDto.UserID,
	}

	// should be updated
	sql := "INSERT INTO btcWallets (address, privateKey, publicKey, wif, pubKeys, scriptType, originalAddress, OAPAddress, createdAt, updatedAt, userId) VALUES (?,?,?,?,?,?,?,?,?,?,?)"
	res := db.QueryRow(sql, wt)
	fmt.Println("sql result is -> ", res)

}

func deleteData() {

}

func selectData() {

}
