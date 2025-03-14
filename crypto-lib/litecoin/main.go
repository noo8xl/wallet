package litecoin

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

	existedAddress := s.db.IsWalletExists(userId, "ltc")
	if !existedAddress {
		return s.generateAddress(userId, 0)
	}

	exceptions.HandleAnHttpExceprion()
	return nil

}

func (s *Service) CreateOneTimeddress(userId int64) *pb.WalletItem {
	return s.generateAddress(userId, 1)
}

func (s *Service) GetBalanceByAddress(address string) *big.Float {

	result, err := s.store.GetAKey(address)
	if val := helpers.BalanceFromStoreFormatter(result, err); val != nil {
		return val
	}

	// ###################################################
	// ######## DOESN"T WORK IN THE TEST-NET! ############
	// ###################################################

	const satoshiPerByte float64 = 0.00000001
	var currentBalance = big.NewFloat(0)

	addressData, err := s.bc.GetAddrBal(address, nil)
	if err != nil {
		exceptions.HandleAnException("<ltc GetAddrBal> got an error: " + err.Error())
	}

	currentBalance = new(big.Float)
	currentBalance.SetPrec(30)
	currentBalance.SetString(addressData.Balance.String())

	bal := new(big.Float).Mul(currentBalance, big.NewFloat(satoshiPerByte))
	s.store.SetAKey(address, bal.String())
	return bal
}

// ===========================================================================================//
// ============================ init connection to the blockchain ============================//
// ================================= and internal functions ==================================//

func (s *Service) generateAddress(userId int64, opt byte) *pb.WalletItem {

	stamp := time.Now().UnixMilli()
	addressKeys, err := s.bc.GenAddrKeychain()

	if err != nil {
		exceptions.HandleAnException("<litecoin generateAddress> got an err: " + err.Error())
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

	wt := &models.LtcWallet{
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
		err := s.db.InsertLtcWalletToPermanent(wt)
		if err != nil {
			exceptions.HandleAnException("<InsertLtcWalletToPermanent> got an error: " + err.Error())
		}
	case 1:
		err := s.db.InsertLtcWalletToOneTimeAddresses(wt)
		if err != nil {
			exceptions.HandleAnException("<InsertLtcWalletToOneTimeAddresses> got an error: " + err.Error())
		}
	default:
		exceptions.HandleAnException(fmt.Sprintf("Unknown opt value %d", opt))
	}

	return &pb.WalletItem{CoinName: "ltc", Address: ""}
}

func initBlockchain() *gobcy.API {

	bc := new(gobcy.API)
	apiToken := config.GetLitecoinAPIKey()

	if apiToken == "" {
		exceptions.HandleAnException("Init <ltc> blockchain got an error. Empty API token.")
	} else {

		bc.Token = apiToken
		bc.Coin = "ltc" // options: "btc","bcy","ltc","doge","eth"
		// bc.Chain = "test3" // depends on coin: "main","test3","test"\
		bc.Chain = "main"
	}

	// log.Println("bc is -> ", bc)

	return bc
}
