package models

// ########################################################################################### //
// ################################ wallet details types ##################################### //
// ########################################################################################### //

type BtcWallet struct {
	ID              int    `json:"id" bson:"id, omitempty"` // auto increment value
	Address         string `json:"address" bson:"address"`
	PrivateKey      string `json:"privateKey" bson:"privateKey"`
	PublicKey       string `json:"publicKey" bson:"publicKey"`
	Wif             string `json:"wif" bson:"wif"`
	PubKeys         string `json:"pubKeys" bson:"pubKeys"`
	ScriptType      string `json:"scriptType" bson:"scriptType"`
	OriginalAddress string `json:"originalAddress" bson:"originalAddress"`
	OAPAddress      string `json:"OAPAddress" bson:"OAPAddress"`
	CreatedAt       int64  `json:"createdAt" bson:"createdAt"` // stamp in ms as value
	UpdatedAt       int64  `json:"updatedAt" bson:"updatedAt"` // stamp in ms as value
	UserID          int64  `json:"userId" bson:"userId"`       // receive via request (telegram chatId)
}

type TonWallet struct {
	ID         int    `json:"id" bson:"id, omitempty"`      // auto increment value
	Address    string `json:"address" bson:"address"`       // address string
	AddrType   int    `json:"addrType" bson:"addrType"`     //
	PrivateKey string `json:"privateKey" bson:"privateKey"` // private
	BitsLen    int    `json:"bitsLen" bson:"bitsLen"`

	// Can be used to operate multiple wallets with the same key and version.
	// use GetSubwallet if you need it.
	Subwallet uint32 `json:"subwallet" bson:"subwallet"`
	CreatedAt int64  `json:"createdAt" bson:"createdAt"` // stamp in ms as value
	UpdatedAt int64  `json:"updatedAt" bson:"updatedAt"` // stamp in ms as value
	UserID    int64  `json:"userId" bson:"userId"`       // receive via request (telegram chatId)
}
