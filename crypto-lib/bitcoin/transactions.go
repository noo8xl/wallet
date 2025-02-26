package bitcoin

import (
	"fmt"
	"math/big"
	"wallet/lib/exceptions"

	pb "wallet/api"

	"github.com/blockcypher/gobcy/v2"
)

// https://www.blockcypher.com/dev/?go#introduction -> doc is here <-

// CreateSingleBitcoinTransactionSkeleton -> create transaction skeleton, sign in locally and send to user for validate it
func (s *BitcoinService) SendSingleBtcTransaction(dto *pb.SendSingleTsxRequest) string {
	var skeleton gobcy.TXSkel
	privateKey := s.db.SelectBtcPrivate(dto.Payee.PeyeeAddress)
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
		exceptions.HandleAnException("<btc create transactions> got an err: " + err.Error())
	}

	// Sign it locally
	err = skeleton.Sign([]string{privateKey})
	if err != nil {
		exceptions.HandleAnException("<btc sign transactions> got an err: " + err.Error())
	}

	// Send TXSkeleton
	skeleton, err = s.bc.SendTX(skeleton)
	if err != nil {
		exceptions.HandleAnException("<btc send transactions> got an err: " + err.Error())
	}

	fmt.Printf("skeleton is => %+v\n", skeleton)

	// save a tsx details to db ?
	skel, err := s.bc.PushTX(skeleton.Trans.Hash)
	if err != nil {
		exceptions.HandleAnException("<btc push transactions> got an err: " + err.Error())
	}
	fmt.Printf("%+v\n", skel)
	return skel.Trans.Hash
}
