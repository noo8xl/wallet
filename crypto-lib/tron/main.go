package tron

import (
	"fmt"
	"io"
	"math/big"
	"net/http"
	"strings"
	"wallet-cli/lib/models"
)

func CreateWallet(userId *string) *models.WalletListItem {
	fmt.Println("generated ton address is\n->", "tmp empty --")
	return &models.WalletListItem{CoinName: "trx", Address: "address will be here"}
}

// GetTrxBalance -> get balance by wallet address
func GetTrxBalance(addr string) *big.Float {

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
		fmt.Println("request err ->", err)
		// err handler message to tg
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))

	// return response.balance / 1000000 // dividing on 1_000_000

	return currentBalance

}
