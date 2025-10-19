package tron

import (
	"fmt"
	"wallet/pkg/exceptions"
	"wallet/wallet/config"
	"wallet/wallet/internal/repository/mysql"
	"wallet/wallet/internal/repository/redis"

	pb "wallet/gen/wallet"
)

type Service struct {
	token   string
	network string
	db      *mysql.DatabaseService
	store   *redis.Store
}

func InitService() *Service {
	svc := initTrxConfig()
	db := mysql.InitDbService()
	s := redis.InitNewStore()
	return &Service{
		token:   svc.token,
		network: svc.network,
		db:      db,
		store:   s,
	}
}

// ===========================================================================================//
// ============================== function for internal usage ================================//
// ===========================================================================================//

func (s *Service) generateAddress(userId int64, opt byte) *pb.WalletItem {

	// TODO: create address *
	// s.db.InsertTonWallet(&models.TonWallet{})

	fmt.Println("userId -> ", userId)

	// key := config.GetAnEncryptionKey()
	// encPrivate, err := helpers.EncryptKey(key, addressKeys.Private)
	// if err != nil {
	// 	exceptions.HandleAnException("private key encoding error")
	// }

	// encPublic, err := helpers.EncryptKey(key, addressKeys.Public)
	// if err != nil {
	// 	exceptions.HandleAnException("public key encoding error")
	// }

	// -> save wallet were to db!
	switch opt {
	case 0:
		// save to permanent addresses

		// if err := s.db.InsertBtcWallet(&wt); err != nil {
		// 	exceptions.HandleAnException("<Database insertion> got an error: " + err.Error())
		// }
	case 1:
		// save to one time addresses

		// if err := s.db.InsertBtcWallet(&wt); err != nil {
		// 	exceptions.HandleAnException("<Database insertion> got an error: " + err.Error())
		// }
	default:
		exceptions.HandleAnException(fmt.Sprintf("Unknown opt value %d", opt))
	}

	return &pb.WalletItem{CoinName: "trx", Address: "address_will_be_here"}
}

func initTrxConfig() *Service {
	var conf = new(Service)
	conf.token = config.GetTronAPIKey()
	conf.network = "https://api.trongrid.io" // mainnet
	// conf.netwotk = "https://api.shasta.trongrid.io" // testnet
	//
	return &Service{token: conf.token, network: conf.network}
}
