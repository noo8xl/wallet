package theopennetwork

import (
	"errors"
	"fmt"
	pb "wallet/api"
)

// DefineaTypeAndSendSingleTransaction -> defina a wallet type (0 - permanenet wallet, 1 - one-time wallet),
// get decoded private key and sign a transaction via it
func (s *Service) DefineaTypeAndSendSingleTransaction(dto *pb.SendSingleTsxRequest) (string, error) {
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
func (s *Service) DefineaTypeAndSendMultipleTransaction(dto *pb.SendMultipleTsxRequest) (string, error) {
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

func (s *Service) sendSingleTransaction(dto *pb.SendSingleTsxRequest, privateKey string) (string, error) {
	fmt.Printf("dto is => %+v\n", dto)
	fmt.Printf("private key  is => %s\n", privateKey)

	return "ton tsx hash", nil
}

func (s *Service) sendMultipleTransactions(dto *pb.SendMultipleTsxRequest, privateKey string) (string, error) {
	fmt.Printf("dto is => %+v\n", dto)
	fmt.Printf("private key  is => %s\n", privateKey)
	return "hash here", nil
}
