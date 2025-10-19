package solana

import (
	"fmt"
	pb "wallet/gen/wallet"
	"wallet/wallet/internal/repository/mysql"
	"wallet/wallet/internal/repository/redis"
)

type Service struct {
	bc    string // tmp
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

//
// ######################################

func (s *Service) generateAddress(userId int64, opt byte) *pb.WalletItem {

	fmt.Println(userId, opt)
	// encrypt the key *

	// // -> save wallet to db!
	// switch opt {
	// case 0:
	// 	err := s.db.InsertEthWalletToPermament(wt)
	// 	if err != nil {
	// 		exceptions.HandleAnException("<InsertEthWalletToPermament> got an error: " + err.Error())
	// 	}
	// case 1:
	// 	err := s.db.InsertEthWalletToOneTimeAddresses(wt)
	// 	if err != nil {
	// 		exceptions.HandleAnException("<InsertEthWalletToOneTimeAddresses> got an error: " + err.Error())
	// 	}
	// default:
	// 	exceptions.HandleAnException(fmt.Sprintf("Unknown opt value %d", opt))
	// }

	return &pb.WalletItem{Address: "", CoinName: "sol"}
}

func initBlockchain() string {
	return ""
}
