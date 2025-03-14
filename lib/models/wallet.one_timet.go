package models

import (
	"crypto/ed25519"

	"github.com/xssnick/tonutils-go/address"
)

type BtcOneTimeWallet struct {
	ID         uint64 `json:"id" bson:"id, omitempty"` // auto increment value
	Address    string `json:"address" bson:"address"`
	PrivateKey string `json:"privateKey" bson:"privateKey"` // encrypted str will be here
	PublicKey  string `json:"publicKey" bson:"publicKey"`   // encrypted str will be here
	Wif        string `json:"wif" bson:"wif"`

	CreatedAt  int64 `json:"createdAt" bson:"createdAt"`   // stamp in ms as value
	UpdatedAt  int64 `json:"updatedAt" bson:"updatedAt"`   // stamp in ms as value
	CustomerId int64 `json:"customerId" bson:"customerId"` // receive via request
	// PubKeys         []string `json:"pubKeys" bson:"pubKeys"`
}

type LtcOneTimeWallet struct {
	ID         uint64 `json:"id" bson:"id, omitempty"` // auto increment value
	Address    string `json:"address" bson:"address"`
	PrivateKey string `json:"privateKey" bson:"privateKey"` // encrypted str will be here
	PublicKey  string `json:"publicKey" bson:"publicKey"`   // encrypted str will be here
	Wif        string `json:"wif" bson:"wif"`

	CreatedAt  int64 `json:"createdAt" bson:"createdAt"`   // stamp in ms as value
	UpdatedAt  int64 `json:"updatedAt" bson:"updatedAt"`   // stamp in ms as value
	CustomerId int64 `json:"customerId" bson:"customerId"` // receive via request
	// PubKeys         []string `json:"pubKeys" bson:"pubKeys"`
}

type DogeOneTimeWallet struct {
	ID         uint64 `json:"id" bson:"id, omitempty"` // auto increment value
	Address    string `json:"address" bson:"address"`
	PrivateKey string `json:"privateKey" bson:"privateKey"` // encrypted str will be here
	PublicKey  string `json:"publicKey" bson:"publicKey"`   // encrypted str will be here
	Wif        string `json:"wif" bson:"wif"`

	CreatedAt  int64 `json:"createdAt" bson:"createdAt"`   // stamp in ms as value
	UpdatedAt  int64 `json:"updatedAt" bson:"updatedAt"`   // stamp in ms as value
	CustomerId int64 `json:"customerId" bson:"customerId"` // receive via request
	// PubKeys         []string `json:"pubKeys" bson:"pubKeys"`
}

type EthOneTimeWallet struct {
	ID         uint64 `json:"id" bson:"id, omitempty"` // auto increment value
	Address    string `json:"address" bson:"address"`
	PrivateKey string `json:"privateKey" bson:"privateKey"` // encrypted str will be here
	PublicKey  string `json:"publicKey" bson:"publicKey"`   // encrypted str will be here
	Wif        string `json:"wif" bson:"wif"`

	CreatedAt  int64 `json:"createdAt" bson:"createdAt"`   // stamp in ms as value
	UpdatedAt  int64 `json:"updatedAt" bson:"updatedAt"`   // stamp in ms as value
	CustomerId int64 `json:"customerId" bson:"customerId"` // receive via request
	// PubKeys         []string `json:"pubKeys" bson:"pubKeys"`
}

type TonOneTimeWallet struct {
	ID         uint64             `json:"id" bson:"id, omitempty"`      // auto increment value
	Address    string             `json:"address" bson:"address"`       // address string
	AddrType   address.AddrType   `json:"addrType" bson:"addrType"`     //
	PrivateKey ed25519.PrivateKey `json:"privateKey" bson:"privateKey"` // private
	BitsLen    uint               `json:"bitsLen" bson:"bitsLen"`

	// Can be used to operate multiple wallets with the same key and version
	// use GetSubwallet if you need it :

	// Subwallet uint32 `json:"subwallet" bson:"subwallet"`
	CreatedAt  int64 `json:"createdAt" bson:"createdAt"`   // stamp in ms as value
	UpdatedAt  int64 `json:"updatedAt" bson:"updatedAt"`   // stamp in ms as value
	CustomerId int64 `json:"customerId" bson:"customerId"` // receive via request
}

type TrxOneTimeWallet struct {
	ID         uint64 `json:"id" bson:"id, omitempty"` // auto increment value
	Address    string `json:"address" bson:"address"`
	PrivateKey string `json:"privateKey" bson:"privateKey"` // encrypted str will be here
	PublicKey  string `json:"publicKey" bson:"publicKey"`   // encrypted str will be here
	Wif        string `json:"wif" bson:"wif"`

	CreatedAt  int64 `json:"createdAt" bson:"createdAt"`   // stamp in ms as value
	UpdatedAt  int64 `json:"updatedAt" bson:"updatedAt"`   // stamp in ms as value
	CustomerId int64 `json:"customerId" bson:"customerId"` // receive via request
}
type SolOneTimeWallet struct {
	ID         uint64 `json:"id" bson:"id, omitempty"` // auto increment value
	Address    string `json:"address" bson:"address"`
	PrivateKey string `json:"privateKey" bson:"privateKey"` // encrypted str will be here
	PublicKey  string `json:"publicKey" bson:"publicKey"`   // encrypted str will be here
	Wif        string `json:"wif" bson:"wif"`

	CreatedAt  int64 `json:"createdAt" bson:"createdAt"`   // stamp in ms as value
	UpdatedAt  int64 `json:"updatedAt" bson:"updatedAt"`   // stamp in ms as value
	CustomerId int64 `json:"customerId" bson:"customerId"` // receive via request
	// PubKeys         []string `json:"pubKeys" bson:"pubKeys"`
}
