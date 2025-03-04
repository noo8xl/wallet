// Package ethereumwallet -> is eth network interacting
package ethereum

import (
	"fmt"
	"math"
	"math/big"
	"time"
	"wallet/config"
	"wallet/database"
	"wallet/lib/exceptions"
	"wallet/lib/models"

	pb "wallet/api"

	"github.com/blockcypher/gobcy/v2"
)

type EthereumService struct {
	bc *gobcy.API
	db *database.DatabaseService
}

func InitEthereumService() *EthereumService {
	bc := initBlockchain()
	db := database.InitDbService()
	return &EthereumService{
		bc: bc,
		db: db,
	}
}

// ----------------------------------------------------------------

// CreateWallet is in charge of creating a new root wallet
func (s *EthereumService) CreatePermanentWallet(userId int64) *pb.WalletItem {

	existedAddress := s.db.IsWalletExists(userId, "eth")
	if !existedAddress {
		return s.generateAddress(userId, 0)
	}
	exceptions.HandleAnHttpExceprion()
	return nil
}

func (s *EthereumService) CreateOneTimeddress(userId int64) *pb.WalletItem {
	return s.generateAddress(userId, 1)
}

// GetEthereumAddressBalance -> get balance by address
func (s *EthereumService) GetEthBalanceByAddress(addr string) *big.Float {
	currentBalance := new(big.Float)

	// ###################################################
	// ######## DOESN"T WORK IN THE TEST-NET! ############
	// ###################################################

	addressData, err := s.bc.GetAddrBal(addr, nil)
	if err != nil {
		exceptions.HandleAnException("<eth GetAddrBal> got an err: " + err.Error())
	}

	currentBalance.SetString(addressData.Balance.String())
	ethValue := new(big.Float).Quo(currentBalance, big.NewFloat(math.Pow10(20)))

	// fmt.Println("balance -> ", currentBalance)
	return ethValue
}

// ===========================================================================================//
// ============================ init the blockchain connection ===============================//
// ===========================================================================================//

func (s *EthereumService) generateAddress(userId int64, opt byte) *pb.WalletItem {

	stamp := time.Now().UnixMilli()

	addressKeys, err := s.bc.GenAddrKeychain()
	if err != nil {
		exceptions.HandleAnException("<eth GenAddrKeychain> got an err: " + err.Error())
	}

	wt := &models.EthWallet{
		Address:         addressKeys.Address,
		PrivateKey:      addressKeys.Private,
		PublicKey:       addressKeys.Public,
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
