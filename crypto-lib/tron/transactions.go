// Package tron -> is tron network interacting
package tron

import (
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"strings"
	"wallet-cli/lib/config"
)

// get api token here
var token = config.GetTronAPIKey()

// https://developers.tron.network/reference/background#note
// -> the doc is here <-

var netwotk = "https://api.trongrid.io" // mainnet
// var netwotk = "https://api.shasta.trongrid.io" // testnet

// ValidateTrxAddress -> validate address in tron network
func ValidateTrxAddress(addr string) bool {

	fmt.Println("token is => ", token)

	// testNet + path + payload
	var response struct {
		Result  bool   `json:"result"`
		Message string `json:"message"`
	}
	path := "/wallet/validateaddress"
	payload := strings.NewReader(strings.Join([]string{"{\"address\":", "\"", addr, "\",", "\"visible\":true}"}, ""))
	url := strings.Join([]string{netwotk, path}, "")

	// fmt.Println("payload ->", payload)
	// fmt.Println(" url -> ", url)

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
	json.Unmarshal(body, &response)
	// fmt.Println(" res is => ", string(body))
	// fmt.Println("response is -> ", response.result)
	return response.Result
}

// CreateWallet -> create wallet in tron network
func CreateWallet(userID string) bool {
	// var err error

	// https://github.com/go-chain/go-tron/blob/master/address/address.go

	// if err != nil {
	// 	notification.SendErrorMessage("Receive an error at create TON wallet.")
	// 	return false
	// }

	return true
}

// GetTrxBalance -> get balance by wallet address
func GetTrxBalance(addr string) *big.Float {

	var currentBalance *big.Float
	var reqString string
	var payload *strings.Reader

	// impl in js ?? <-----------------------

	// reqUrl := strings.Join([]string{"https://apilist.tronscan.org/api/account?address=", addr}, "")
	// const payload: any = { "address": this.address }
	// const coinData: any = await axios({
	// 	method: 'GET',
	// 	url: reqUrl,
	// 	headers: payload
	// })
	// console.log('coinData => ', coinData);
	// const dataArray: any = coinData.data.tokens
	// let trxNetworkCoinBalance: number;
	// console.log('received data  => ', coinData.data.tokens);
	// for (let i = 0; i < dataArray.length; i++)
	// trxNetworkCoinBalance = +dataArray[i].balance / 1_000_000

	// var response struct {
	// 	balance         int64
	// 	blockIdentifier struct {
	// 		hash   string
	// 		number int64
	// 	}
	// }

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

// ===========================================================================================//
// ============================== function for internal usage ================================//
// ===========================================================================================//

// func generateTrxAddress() {

// }

// func initTrxNetwork() {

// }
