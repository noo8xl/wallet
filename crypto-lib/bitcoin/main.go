// Package bitcoinwallet -> is all of btc network interacting
package bitcoin

import (
	"fmt"
	"math/big"
	"time"
	"wallet/config"
	"wallet/database"
	"wallet/lib/cache"
	"wallet/lib/exceptions"
	"wallet/lib/helpers"
	"wallet/lib/models"

	// "wallet/lib/models"

	pb "wallet/api"

	"github.com/blockcypher/gobcy/v2"
)

type Service struct {
	bc    *gobcy.API
	db    *database.DatabaseService
	store *cache.Store
}

func InitService() *Service {
	bc := initBlockchain()
	db := database.InitDbService()
	s := cache.InitNewStore()
	return &Service{
		bc:    bc,
		db:    db,
		store: s,
	}
}

// CreateWallet is in charge of creating a new root wallet
func (s *Service) CreatePermanentWallet(userId int64) *pb.WalletItem {
	existedAddress, err := s.db.IsWalletExists(userId, "btc", 0)
	if err == nil {
		if !existedAddress {
			return s.generateAddress(userId, 0)
		}
	}

	return nil
}

func (s *Service) CreateOneTimeddress(userId int64) *pb.WalletItem {
	return s.generateAddress(userId, 1)
}

// GetBitcoinAddressBalance -> get balance by address
func (s *Service) GetBalanceByAddress(address string) (*big.Float, error) {

	result, err := s.store.GetAKey(address)
	if err != nil {
		return nil, err
	}
	if val := helpers.BalanceFromStoreFormatter(result, err); val != nil {
		return val, nil
	}
	// log.Println("adr -> ", address)
	// ###################################################
	// ######## DOESN"T WORK IN THE TEST-NET! ############
	// ###################################################

	// balance value from API will be receive in satoshi value
	// to calculate it in btc value -> should multiply on 0.00_000_001
	// https://buybitcoinworldwide.com/satoshi-usd/

	const satoshiPerByte float64 = 0.00000001
	var currentBalance = big.NewFloat(0)

	addressData, err := s.bc.GetAddrBal(address, nil)
	if err != nil {
		exceptions.HandleAnException("<btc GetAddrBal> got an error: " + err.Error())
	}

	// fmt.Println("addressData -> ")
	// helpers.PrintPretty(addressData)
	// fmt.Println("addressData.Balance -> ", addressData.Balance.String())

	currentBalance = new(big.Float)
	currentBalance.SetPrec(30)
	currentBalance.SetString(addressData.Balance.String())

	bal := new(big.Float).Mul(currentBalance, big.NewFloat(satoshiPerByte))
	s.store.SetAKey(address, bal.String())
	return bal, nil
}

// ===========================================================================================//
// ============================ init connection to the blockchain ============================//
// ================================= and internal functions ==================================//

func (s *Service) generateAddress(userId int64, opt byte) *pb.WalletItem {

	stamp := time.Now().UnixMilli()
	addressKeys, err := s.bc.GenAddrKeychain()

	if err != nil {
		exceptions.HandleAnException("<btc generateAddress> got an err: " + err.Error())
	}

	key := config.GetAnEncryptionKey()
	encPrivate, err := helpers.EncryptKey(key, addressKeys.Private)
	if err != nil {
		exceptions.HandleAnException("private key encoding error")
	}

	encPublic, err := helpers.EncryptKey(key, addressKeys.Public)
	if err != nil {
		exceptions.HandleAnException("public key encoding error")
	}

	wt := &models.BtcWallet{
		Address:         addressKeys.Address,
		PrivateKey:      encPrivate,
		PublicKey:       encPublic,
		Wif:             addressKeys.Wif,
		ScriptType:      addressKeys.ScriptType,
		OriginalAddress: addressKeys.OriginalAddress,
		OAPAddress:      addressKeys.OAPAddress,
		CreatedAt:       stamp,
		UpdatedAt:       stamp,
		CustomerId:      userId,
		// PubKeys:         addressKeys.PubKeys,
	}

	// -> save wallet to db!
	switch opt {
	case 0:
		err := s.db.InsertBtcWalletToPermanent(wt)
		if err != nil {
			exceptions.HandleAnException("<InsertBtcWalletToPermanent> got an error: " + err.Error())
		}
	case 1:
		err := s.db.InsertBtcWalletToOneTimeAddresses(wt)
		if err != nil {
			exceptions.HandleAnException("<InsertBtcWalletToOneTimeAddresses> got an error: " + err.Error())
		}
	default:
		exceptions.HandleAnException(fmt.Sprintf("Unknown opt value %d", opt))
	}

	return &pb.WalletItem{CoinName: "btc", Address: wt.Address}
}

func initBlockchain() *gobcy.API {

	bc := new(gobcy.API)
	apiToken := config.GetBitcoinAPIKey()

	if apiToken == "" {
		exceptions.HandleAnException("Init <btc> blockchain got an error. Empty API token.")
	} else {

		bc.Token = apiToken
		bc.Coin = "btc"    // options: "btc","bcy","ltc","doge","eth"
		bc.Chain = "test3" // depend on coin: "main","test3","test"\
		// bc.Chain = "main"
	}

	// log.Println("bc is -> ", bc)

	return bc
}
