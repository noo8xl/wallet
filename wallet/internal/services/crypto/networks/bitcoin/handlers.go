package bitcoin

import (
	"math/big"
	pb "wallet/gen/wallet"
	"wallet/pkg/exceptions"
	"wallet/wallet/pkg/utils"
)

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

func (s *Service) CreateOneTimeAddress(userId int64) *pb.WalletItem {
	return s.generateAddress(userId, 1)
}

// GetBitcoinAddressBalance -> get balance by address
func (s *Service) GetBalanceByAddress(address string) (*big.Float, error) {

	result, err := s.store.GetAKey(address)
	if err != nil {
		return nil, err
	}
	if val := utils.BalanceFromStoreFormatter(result, err); val != nil {
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
