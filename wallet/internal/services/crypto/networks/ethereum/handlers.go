package ethereum

import (
	"math"
	"math/big"
	pb "wallet/gen/wallet"
	"wallet/wallet/pkg/utils"
)

// CreateWallet is in charge of creating a new root wallet
func (s *Service) CreatePermanentWallet(userId int64) *pb.WalletItem {
	existedAddress, err := s.db.IsWalletExists(userId, "eth", 0)
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

// GetEthereumAddressBalance -> get balance by address
func (s *Service) GetBalanceByAddress(address string) (*big.Float, error) {

	result, err := s.store.GetAKey(address)
	if err != nil {
		return nil, err
	}
	if val := utils.BalanceFromStoreFormatter(result, err); val != nil {
		return val, nil
	}

	currentBalance := new(big.Float)

	// ###################################################
	// ######## DOESN"T WORK IN THE TEST-NET! ############
	// ###################################################

	addressData, err := s.bc.GetAddrBal(address, nil)
	if err != nil {
		return nil, err
	}

	currentBalance.SetString(addressData.Balance.String())
	ethValue := new(big.Float).Quo(currentBalance, big.NewFloat(math.Pow10(20)))

	// fmt.Println("balance -> ", currentBalance)
	s.store.SetAKey(address, ethValue.String())
	return ethValue, nil
}
