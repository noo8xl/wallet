package tron

import (
	"fmt"
	"io"
	"math/big"
	"net/http"
	"strings"
	pb "wallet/gen/wallet"
	"wallet/wallet/pkg/utils"
)

func (s *Service) CreatePermanentWallet(userId int64) *pb.WalletItem {
	existedAddress, err := s.db.IsWalletExists(userId, "trx", 0)
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

// GetTrxBalance -> get balance by wallet address
func (s *Service) GetBalanceByAddress(addr string) (*big.Float, error) {

	result, err := s.store.GetAKey(addr)
	if err != nil {
		return nil, err
	}
	if val := utils.BalanceFromStoreFormatter(result, err); val != nil {
		return val, nil
	}

	var currentBalance *big.Float
	var reqString string
	var payload *strings.Reader

	currentBalance = new(big.Float)
	currentBalance.SetString("123")

	fmt.Println(addr)
	url := "https://api.shasta.trongrid.io/wallet/getaccountbalance"
	reqString = strings.Join([]string{"{\"account_identifier\":{\"address\":", "\"", addr, "\"}"}, "")

	payload = strings.NewReader("{\"account_identifier\":{\"address\":\"TFH9ufxh8CpYxa7W9LdA8Y1iHHptYErvb7\"}")

	fmt.Println("payload =>\n", payload)
	fmt.Println("payload =>\n", reqString)

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))

	// return response.balance / 1000000 // dividing on 1_000_000
	s.store.SetAKey(addr, string(body))
	return currentBalance, nil

}
