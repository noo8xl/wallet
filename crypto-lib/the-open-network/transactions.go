package theopennetwork

import (
	"log"
	pb "wallet/api"
	"wallet/config"
	"wallet/lib/exceptions"
	"wallet/lib/helpers"
)

func (s *Service) SendSingleTransaction(dto *pb.SendSingleTsxRequest) string {
	// manage smth and send trx ->

	key := []byte(config.GetAnEncryptionKey())
	encryptedPk := s.db.SelectTonPrivate(dto.Payee.PeyeeAddress)
	privateKey, err := helpers.DecryptByteArr(key, string(encryptedPk))
	if err != nil {
		exceptions.HandleAnException("SendSingleTonTransaction got an error: " + err.Error())
	}

	log.Printf("ton private key is -> %s", privateKey)

	return "ton tsx hash"
}

func (s *Service) SendMultipleTransactions(dto *pb.SendMultipleTsxRequest) string {
	return "hash here"
}
