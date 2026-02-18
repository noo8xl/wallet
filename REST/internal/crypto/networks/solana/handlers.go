package solana

import (
	"math/big"
	pb "wallet/gen/wallet"
)

func (s *Service) CreatePermanentWallet(userId int64) *pb.WalletItem {
	existedAddress, err := s.db.IsWalletExists(userId, "sol", 0)
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

	bal := big.NewFloat(0)

	return bal, nil
}
