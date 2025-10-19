// Package models -> is package for describe all models, interfaces, data types, schemas, etc.
package models

import "math/big"

// -> API response types -<

type SendTransactionDto struct {
	SenderAddress    string `json:"senderAddress" xml:"senderAddress" form:"senderAddress"`
	CoinName         string `json:"coinName" xml:"coinName" form:"coinName"`
	RecipientAddress string `json:"recipientAddress" xml:"recipientAddress" form:"recipientAddress"`
	Amount           string `json:"amount" xml:"amount" form:"amount"`
}

// ResponseBalance -> is a get balance cmd response
type ResponseBalance struct {
	CoinName     string
	CoinBalance  *big.Float
	CurrencyType string
	FiatAmount   *big.Float
}

// WalletListItem -> is response of create wallet func
type WalletListItem struct {
	CoinName string
	Address  string
}
