package tron

import (
	"fmt"
	"io"
	"math/big"
	"net/http"
	"strings"
)

func CreateWallet(userId string) string {
	return ""
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
