// Package ethereumwallet -> is eth network interacting
package ethereum

import (
	"fmt"
	"math"
	"math/big"
	"time"
	"wallet/config"
	"wallet/database"
	"wallet/lib/cache"
	"wallet/lib/exceptions"
	"wallet/lib/helpers"
	"wallet/lib/models"

	pb "wallet/api"

	"github.com/blockcypher/gobcy/v2"
)

type Service struct {
	bc    *gobcy.API
	db    *database.DatabaseService
	store *cache.Store
}

func InitEthereumService() *Service {
	bc := initBlockchain()
	db := database.InitDbService()
	s := cache.InitNewStore()
	return &Service{
		bc:    bc,
		db:    db,
		store: s,
	}
}

// ----------------------------------------------------------------

// CreateWallet is in charge of creating a new root wallet
func (s *Service) CreatePermanentWallet(userId int64) *pb.WalletItem {

	existedAddress := s.db.IsWalletExists(userId, "eth")
	if !existedAddress {
		return s.generateAddress(userId, 0)
	}
	exceptions.HandleAnHttpExceprion()
	return nil
}

func (s *Service) CreateOneTimeddress(userId int64) *pb.WalletItem {
	return s.generateAddress(userId, 1)
}

// GetEthereumAddressBalance -> get balance by address
func (s *Service) GetBalanceByAddress(address string) *big.Float {

	result, err := s.store.GetAKey(address)
	if val := helpers.BalanceFromStoreFormatter(result, err); val != nil {
		return val
	}

	currentBalance := new(big.Float)

	// ###################################################
	// ######## DOESN"T WORK IN THE TEST-NET! ############
	// ###################################################

	addressData, err := s.bc.GetAddrBal(address, nil)
	if err != nil {
		exceptions.HandleAnException("<eth GetAddrBal> got an err: " + err.Error())
	}

	currentBalance.SetString(addressData.Balance.String())
	ethValue := new(big.Float).Quo(currentBalance, big.NewFloat(math.Pow10(20)))

	// fmt.Println("balance -> ", currentBalance)
	s.store.SetAKey(address, ethValue.String())
	return ethValue
}

// ===========================================================================================//
// ============================ init the blockchain connection ===============================//
// ===========================================================================================//

func (s *Service) generateAddress(userId int64, opt byte) *pb.WalletItem {

	stamp := time.Now().UnixMilli()

	addressKeys, err := s.bc.GenAddrKeychain()
	if err != nil {
		exceptions.HandleAnException("<eth GenAddrKeychain> got an err: " + err.Error())
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

	wt := &models.EthWallet{
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
		err := s.db.InsertEthWalletToPermament(wt)
		if err != nil {
			exceptions.HandleAnException("<InsertEthWalletToPermament> got an error: " + err.Error())
		}
	case 1:
		err := s.db.InsertEthWalletToOneTimeAddresses(wt)
		if err != nil {
			exceptions.HandleAnException("<InsertEthWalletToOneTimeAddresses> got an error: " + err.Error())
		}
	default:
		exceptions.HandleAnException(fmt.Sprintf("Unknown opt value %d", opt))
	}

	return &pb.WalletItem{Address: wt.Address, CoinName: "eth"}
}

func initBlockchain() *gobcy.API {

	bc := new(gobcy.API)
	apiToken := config.GetBitcoinAPIKey()

	if apiToken == "" {
		exceptions.HandleAnException("Init <eth> blockchain got an error. Empty API token.")
	} else {

		bc.Token = apiToken
		bc.Coin = "eth" // options: "btc","bcy","ltc","doge","eth"
		// bc.Chain = "test3" // depend on coin: "main","test3","test"\
		bc.Chain = "main"
	}

	return bc
}
