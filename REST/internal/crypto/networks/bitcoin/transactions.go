package bitcoin

import (
	"errors"
	"fmt"
	"math/big"

	pb "wallet/gen/wallet"

	"github.com/blockcypher/gobcy/v2"
)

// DefineaTypeAndSendSingleTransaction -> defina a wallet type (0 - permanenet wallet, 1 - one-time wallet),
// get decoded private key and sign a transaction via it
func (s *Service) DefineBlockchainAndSendSingleTransaction(dto *pb.SendSingleTsxRequest) (string, error) {
	var hash string
	var err error
	var privateKey string

	switch dto.WalletType {
	case 0:
		privateKey, err = s.db.DefineaBlockchainAndGetPrivateKey(dto.Payee.PeyeeAddress, "btc", byte(dto.WalletType))
		if err != nil {
			return "", err
		}

		hash, err = s.sendSingleTransaction(dto, privateKey)
	case 1:
		privateKey, err = s.db.DefineaBlockchainAndGetPrivateKey(dto.Payee.PeyeeAddress, "btc", byte(dto.WalletType))
		if err != nil {
			return "", err
		}

		hash, err = s.sendSingleTransaction(dto, privateKey)
	default:
		return "", errors.New("invalid option")
	}

	return hash, err
}

// DefineaTypeAndSendMultipleTransaction -> defina a wallet type (0 - permanenet wallet, 1 - one-time wallet),
// get decoded private key and sign a transaction via it
func (s *Service) DefineBlockchainAndSendMultipleTransaction(dto *pb.SendMultipleTsxRequest) (string, error) {
	var hash string
	var err error
	var privateKey string

	switch dto.WalletType {
	case 0:
		privateKey, err = s.db.DefineaBlockchainAndGetPrivateKey(dto.Payee.PeyeeAddress, "btc", byte(dto.WalletType))
		if err != nil {
			return "", err
		}

		hash, err = s.sendMultipleTransactions(dto, privateKey)
	case 1:
		privateKey, err = s.db.DefineaBlockchainAndGetPrivateKey(dto.Payee.PeyeeAddress, "btc", byte(dto.WalletType))
		if err != nil {
			return "", err
		}

		hash, err = s.sendMultipleTransactions(dto, privateKey)
	default:
		return "", errors.New("invalid option")
	}

	return hash, err
}

// #################################################################################################################
// #################################################################################################################
// #################################################################################################################

// CreateSingleBitcoinTransactionSkeleton -> create transaction skeleton, sign in locally and send to user for validate it
func (s *Service) sendSingleTransaction(dto *pb.SendSingleTsxRequest, privateKey string) (string, error) {
	var skeleton gobcy.TXSkel

	amount := new(big.Int)
	amount.SetString(dto.Beneficiar.Amount, 10)

	//Post New TXSkeleton
	skeleton, err := s.bc.NewTX(gobcy.TempNewTX(dto.Payee.PeyeeAddress, dto.Beneficiar.BeneficiarAddress, *amount), false)
	if err != nil {
		return "", err
	}

	// Sign it locally
	err = skeleton.Sign([]string{privateKey})
	if err != nil {
		return "", err
	}

	// Send TXSkeleton
	skeleton, err = s.bc.SendTX(skeleton)
	if err != nil {
		return "", err
	}

	fmt.Printf("skeleton is => %+v\n", skeleton)

	// save a tsx details to db ?
	skel, err := s.bc.PushTX(skeleton.Trans.Hash)
	if err != nil {
		return "", err
	}
	// fmt.Printf("%+v\n", skel)
	return skel.Trans.Hash, nil
}

func (s *Service) sendMultipleTransactions(dto *pb.SendMultipleTsxRequest, privateKey string) (string, error) {
	fmt.Printf("dto is => %+v\n", dto)
	fmt.Printf("private key  is => %s\n", privateKey)
	return "hash here", nil
}
