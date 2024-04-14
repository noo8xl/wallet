package models

import (
	"crypto/ed25519"

	"github.com/xssnick/tonutils-go/address"
)

// ########################################################################################### //
// ################################ wallet details types ##################################### //
// ########################################################################################### //

type BtcWallet struct {
	ID         int    `json:"id" bson:"id, omitempty"` // auto increment value
	Address    string `json:"address" bson:"address"`
	PrivateKey string `json:"privateKey" bson:"privateKey"`
	PublicKey  string `json:"publicKey" bson:"publicKey"`
	Wif        string `json:"wif" bson:"wif"`
	// PubKeys         []string `json:"pubKeys" bson:"pubKeys"`
	ScriptType      string `json:"scriptType" bson:"scriptType"`
	OriginalAddress string `json:"originalAddress" bson:"originalAddress"`
	OAPAddress      string `json:"OAPAddress" bson:"OAPAddress"`
	CreatedAt       int64  `json:"createdAt" bson:"createdAt"` // stamp in ms as value
	UpdatedAt       int64  `json:"updatedAt" bson:"updatedAt"` // stamp in ms as value
	UserId          string `json:"userId" bson:"userId"`       // receive via request
}

// create table btcWallets (id int auto_increment primary key, address varchar(100) not null, privateKey varchar(500) not null, publicKey varchar(500) not null, wif varchar(500) not null, scriptType varchar(500), originalAddress varchar(500), oapAddress varchar(500), createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP, updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP, userId varchar(200) not null );

type TonWallet struct {
	ID         int                `json:"id" bson:"id, omitempty"`      // auto increment value
	Address    string             `json:"address" bson:"address"`       // address string
	AddrType   address.AddrType   `json:"addrType" bson:"addrType"`     //
	PrivateKey ed25519.PrivateKey `json:"privateKey" bson:"privateKey"` // private
	BitsLen    uint               `json:"bitsLen" bson:"bitsLen"`

	// Can be used to operate multiple wallets with the same key and version.
	// use GetSubwallet if you need it.
	// Subwallet uint32 `json:"subwallet" bson:"subwallet"`
	CreatedAt int64  `json:"createdAt" bson:"createdAt"` // stamp in ms as value
	UpdatedAt int64  `json:"updatedAt" bson:"updatedAt"` // stamp in ms as value
	UserId    string `json:"userId" bson:"userId"`       // receive via request
}
