package solana

import (
	"fmt"
	pb "wallet/api"
	"wallet/database"
	"wallet/lib/cache"
	"wallet/lib/exceptions"
)

type Service struct {
	bc    string // tmp
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

func (s *Service) CreatePermanentWallet(userId int64) *pb.WalletItem {
	existedAddress := s.db.IsWalletExists(userId, "sol")
	if !existedAddress {
		return s.generateAddress(userId, 0)
	}
	exceptions.HandleAnHttpExceprion()
	return nil
}

func (s *Service) CreateOneTimeddress(userId int64) *pb.WalletItem {
	return s.generateAddress(userId, 1)
}

func (s *Service) GetBalanceByAddress(address string) string {
	return "balance will be here"
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
