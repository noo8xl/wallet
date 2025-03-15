// Package tron -> is tron network interacting
package tron

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"wallet/lib/exceptions"

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
		privateKey, err = s.db.DefineaBlockchainAndGetPrivateKey(dto.Payee.PeyeeAddress, "trx", byte(dto.WalletType))
		if err != nil {
			return "", err
		}

		hash, err = s.sendSingleTransaction(dto, privateKey)
	case 1:
		privateKey, err = s.db.DefineaBlockchainAndGetPrivateKey(dto.Payee.PeyeeAddress, "trx", byte(dto.WalletType))
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
		privateKey, err = s.db.DefineaBlockchainAndGetPrivateKey(dto.Payee.PeyeeAddress, "trx", byte(dto.WalletType))
		if err != nil {
			return "", err
		}

		hash, err = s.sendMultipleTransactions(dto, privateKey)
	case 1:
		privateKey, err = s.db.DefineaBlockchainAndGetPrivateKey(dto.Payee.PeyeeAddress, "trx", byte(dto.WalletType))
		if err != nil {
			return "", err
		}

		hash, err = s.sendMultipleTransactions(dto, privateKey)
	default:
		return "", errors.New("invalid option")
	}

	return hash, err
}

// https://developers.tron.network/reference/background#note
// -> the doc is here <-
//
// ValidateTrxAddress -> validate address in tron network
func (s *Service) ValidateTrxAddress(addr string) bool {

	// testNet + path + payload
	var response struct {
		Result  bool   `json:"result"`
		Message string `json:"message"`
	}
	path := "/wallet/validateaddress"
	payload := strings.NewReader(strings.Join([]string{"{\"address\":", "\"", addr, "\",", "\"visible\":true}"}, ""))
	url := strings.Join([]string{s.network, path}, "")

	// fmt.Println("payload ->", payload)
	// fmt.Println(" url -> ", url)

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		exceptions.HandleAnException("<trx validate address> got an http err: " + err.Error())

	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &response)
	// fmt.Println(" res is => ", string(body))
	// fmt.Println("response is -> ", response.result)
	return response.Result
}

func (s *Service) sendSingleTransaction(dto *pb.SendSingleTsxRequest, privateKey string) (string, error) {

	fmt.Printf("dto is => %+v\n", dto)
	fmt.Printf("private key  is => %s\n", privateKey)
	// save a tsx details to db
	return "trx tsx hash", nil
}

func (s *Service) sendMultipleTransactions(dto *pb.SendMultipleTsxRequest, privateKey string) (string, error) {
	fmt.Printf("dto is => %+v\n", dto)
	fmt.Printf("private key  is => %s\n", privateKey)
	return "hash here", nil
}
