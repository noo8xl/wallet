package theopennetwork

import (
	"context"
	"fmt"
	"math/big"
	pb "wallet/gen/wallet"
	"wallet/wallet/pkg/utils"

	"github.com/xssnick/tonutils-go/address"
)

func (s *Service) CreatePermanentWallet(userId int64) *pb.WalletItem {
	existedAddress, err := s.db.IsWalletExists(userId, "ton", 0)
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

// GetTonBalanceByAddress -> get balance value by coin address
func (s *Service) GetBalanceByAddress(a string) (*big.Float, error) {

	result, err := s.store.GetAKey(a)
	if err != nil {
		return nil, err
	}
	if val := utils.BalanceFromStoreFormatter(result, err); val != nil {
		return val, nil
	}

	curBalance := new(big.Float)
	ctx := context.Background()
	addr := address.MustParseAddr(a)

	// we need fresh block info to run get methods
	blcn, err := s.client.CurrentMasterchainInfo(ctx)
	if err != nil {
		return nil, err
	}

	// we use WaitForBlock to make sure block is ready,
	// it is optional but escapes us from liteserver block not ready errors
	acc, err := s.client.WaitForBlock(blcn.SeqNo).GetAccount(ctx, blcn, addr)
	if err != nil {
		return nil, err
	}

	fmt.Println("acc info =>")
	utils.PrintPretty(blcn)

	fmt.Printf("Is active: %v\n", acc.IsActive)
	if !acc.IsActive {
		return nil, err
	}

	fmt.Printf("Status: %s\n", acc.State.Status)

	curBalance.SetString(acc.State.Balance.String())
	// fmt.Println(" curBalance --> ", curBalance)
	s.store.SetAKey(a, curBalance.String())
	return curBalance, nil
}
