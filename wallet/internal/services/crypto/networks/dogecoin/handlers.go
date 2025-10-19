package dogecoin

import (
	"math/big"
	pb "wallet/gen/wallet"
	"wallet/wallet/pkg/utils"
)

// CreateWallet is in charge of creating a new root wallet
func (s *Service) CreatePermanentWallet(userId int64) *pb.WalletItem {
	existedAddress, err := s.db.IsWalletExists(userId, "doge", 0)
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

func (s *Service) GetBalanceByAddress(address string) (*big.Float, error) {

	result, err := s.store.GetAKey(address)
	if err != nil {
		return nil, err
	}
	if val := utils.BalanceFromStoreFormatter(result, err); val != nil {
		return val, nil
	}

	// ###################################################
	// ######## DOESN"T WORK IN THE TEST-NET! ############
	// ###################################################

	const satoshiPerByte float64 = 0.00000001
	var currentBalance = big.NewFloat(0)

	addressData, err := s.bc.GetAddrBal(address, nil)
	if err != nil {
		return nil, err
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
