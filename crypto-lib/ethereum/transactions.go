package ethereum

import (
	"errors"
	"fmt"
	"math/big"
	"wallet/lib/exceptions"

	pb "wallet/api"

	"github.com/blockcypher/gobcy/v2"
)

// DefineaTypeAndSendSingleTransaction -> defina a wallet type (0 - permanenet wallet, 1 - one-time wallet),
// get decoded private key and sign a transaction via it
func (s *Service) DefineaTypeAndSendSingleTransaction(dto *pb.SendSingleTsxRequest) (string, error) {
	var hash string
	var err error
	var privateKey string

	switch dto.WalletType {
	case 0:
		privateKey, err = s.db.DefineaBlockchainAndGetPrivateKey(dto.Payee.PeyeeAddress, "eth", byte(dto.WalletType))
		if err != nil {
			return "", err
		}

		hash, err = s.sendSingleTransaction(dto, privateKey)
	case 1:
		privateKey, err = s.db.DefineaBlockchainAndGetPrivateKey(dto.Payee.PeyeeAddress, "eth", byte(dto.WalletType))
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
func (s *Service) DefineaTypeAndSendMultipleTransaction(dto *pb.SendMultipleTsxRequest) (string, error) {
	var hash string
	var err error
	var privateKey string

	switch dto.WalletType {
	case 0:
		privateKey, err = s.db.DefineaBlockchainAndGetPrivateKey(dto.Payee.PeyeeAddress, "eth", byte(dto.WalletType))
		if err != nil {
			return "", err
		}

		hash, err = s.sendMultipleTransactions(dto, privateKey)
	case 1:
		privateKey, err = s.db.DefineaBlockchainAndGetPrivateKey(dto.Payee.PeyeeAddress, "eth", byte(dto.WalletType))
		if err != nil {
			return "", err
		}

		hash, err = s.sendMultipleTransactions(dto, privateKey)
	default:
		return "", errors.New("invalid option")
	}

	return hash, err
}

// ###########################################################################################
// ###########################################################################################
// ###########################################################################################

func (s *Service) sendSingleTransaction(dto *pb.SendSingleTsxRequest, privateKey string) (string, error) {
	var skeleton gobcy.TXSkel
	var err error
	amount := new(big.Int)
	amount.SetString(dto.Beneficiar.Amount, 10)

	// use faucet to fund first
	// _, err = bc.Faucet(addressFrom, 3e5)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	//Post New TXSkeleton
	skeleton, err = s.bc.NewTX(gobcy.TempNewTX(dto.Payee.PeyeeAddress, dto.Beneficiar.BeneficiarAddress, *amount), false)
	if err != nil {
		exceptions.HandleAnException("<eth create transactions> got an err: " + err.Error())
	}

	// Sign it locally
	err = skeleton.Sign([]string{privateKey})
	if err != nil {
		exceptions.HandleAnException("<eth sign transactions> got an err: " + err.Error())
	}

	// Send TXSkeleton
	skeleton, err = s.bc.SendTX(skeleton)
	if err != nil {
		exceptions.HandleAnException("<eth send transactions> got an err: " + err.Error())
	}
	fmt.Printf("skeleton is => %+v\n", skeleton)

	// save a tsx details to db
	skel, err := s.bc.PushTX(skeleton.Trans.Hash)
	if err != nil {
		exceptions.HandleAnException("<eth push transactions> got an err: " + err.Error())
	}
	fmt.Printf("%+v\n", skel)
	return skel.Trans.Hash, nil
}

func (s *Service) sendMultipleTransactions(dto *pb.SendMultipleTsxRequest, privateKey string) (string, error) {
	fmt.Printf("dto is => %+v\n", dto)
	fmt.Printf("private key  is => %s\n", privateKey)
	return "hash here", nil
}
