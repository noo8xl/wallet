// Package models -> is package for describe all models, interfaces, data types, schemas, etc.
package models

// -> API response types below -<

type WalletItemResponse struct {
	UserID      uint64
	Address     string
	CoinName    string
	CoinBalance string  // *big.Float => balance.Text('f', -1)
	FiatValue   float64 // a balance in retreived currency type
	// update? *
}

type SendTransactionDto struct {
	SenderAddress    string `json:"senderAddress" xml:"senderAddress" form:"senderAddress"`
	CoinName         string `json:"coinName" xml:"coinName" form:"coinName"`
	RecipientAddress string `json:"recipientAddress" xml:"recipientAddress" form:"recipientAddress"`
	Amount           string `json:"amount" xml:"amount" form:"amount"`
}

// WalletStat -> for the mongoDB for fast access from client API
// without ctypto API call
type WalletStats struct {
	ID           string  `bson:"_id,omitempty"`
	UserEmail    string  `json:"userEmail" bson:"userEmail"`       // receive via request
	UserId       int64   `json:"userId" bson:"userId"`             // receive via request
	TotalBalance float64 `json:"totalBalance" bson:"totalBalance"` // default 0 -> total user wallet balance (float value 0.00) displayed in chosen fiat currency
	CurrencyType string  `json:"currencyType" bson:"currencyType"` // default USD -> ["AUD","USD", "EUR", "AED", "RUB"]
	UpdatedAt    int64   `json:"updatedAt" bson:"updatedAt"`       // update once a 300 ms OR if manipulating wallet by user <-
}

// WalletCoinItem- > describe wallet item types
type WalletCoinItem struct {
	UserId      int64   `json:"userId" bson:"userId"` // receive via request
	CoinName    string  `json:"coinName" bson:"coinName"`
	Address     string  `json:"address" bson:"address"`
	CoinBalance float64 `json:"coinBalance" bson:"coinBalance"` // as float (0.00)
	FiatBalance float32 `json:"fiatBalance" bson:"fiatBalance"` // in chosen currency type ["AUD","USD", "EUR", "AED", "RUB"]
	WalletId    string  `json:"walletId" bson:"walletId"`       // ref to the WalletStat _id ->
}

type ParserCache struct {
	Status    string // current status
	UpdatedAt int64  // last status update
	StartFrom int64  // date to parse from
}
