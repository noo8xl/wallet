package ethereum

import (
	"fmt"
	"math/big"
	"wallet/config"
	"wallet/lib/exceptions"
	"wallet/lib/helpers"

	pb "wallet/api"

	"github.com/blockcypher/gobcy/v2"
)

func (s *Service) SendSingleTransaction(dto *pb.SendSingleTsxRequest) string {
	var skeleton gobcy.TXSkel

	key := config.GetAnEncryptionKey()
	encryptedPk := s.db.SelectEthPrivate(dto.Payee.PeyeeAddress)
	privateKey, err := helpers.DecryptKey(key, encryptedPk)
	if err != nil {
		exceptions.HandleAnException("SendSingleEthTransaction got an error: " + err.Error())
	}

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
	return skel.Trans.Hash
}

func (s *Service) SendMultipleTransactions(dto *pb.SendMultipleTsxRequest) string {
	return "hash here"
}
