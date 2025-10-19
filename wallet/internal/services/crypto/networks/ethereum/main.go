// Package ethereumwallet -> is eth network interacting
package ethereum

import (
	"fmt"
	"time"
	"wallet/pkg/exceptions"
	"wallet/wallet/config"
	"wallet/wallet/internal/repository/mysql"
	"wallet/wallet/internal/repository/redis"
	models "wallet/wallet/pkg/models/wallet"
	"wallet/wallet/pkg/utils"

	pb "wallet/gen/wallet"

	"github.com/blockcypher/gobcy/v2"
)

type Service struct {
	bc    *gobcy.API
	db    *mysql.DatabaseService
	store *redis.Store
}

func InitService() *Service {
	bc := initBlockchain()
	db := mysql.InitDbService()
	s := redis.InitNewStore()
	return &Service{
		bc:    bc,
		db:    db,
		store: s,
	}
}

// ----------------------------------------------------------------

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
	encPrivate, err := utils.EncryptKey(key, addressKeys.Private)
	if err != nil {
		exceptions.HandleAnException("private key encoding error")
	}

	encPublic, err := utils.EncryptKey(key, addressKeys.Public)
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
