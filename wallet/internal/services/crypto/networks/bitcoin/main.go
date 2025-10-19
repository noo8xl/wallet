// Package bitcoinwallet -> is all of btc network interacting
package bitcoin

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
	encPrivate, err := utils.EncryptKey(key, addressKeys.Private)
	if err != nil {
		exceptions.HandleAnException("private key encoding error")
	}

	encPublic, err := utils.EncryptKey(key, addressKeys.Public)
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
