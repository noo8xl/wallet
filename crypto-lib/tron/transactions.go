// Package tron -> is tron network interacting
package tron

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"wallet-cli/config"
	"wallet-cli/lib/models"
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

func SendSingleTrxTransaction(dto models.SendTransactionDto) {
	// return
}

// ===========================================================================================//
// ============================== function for internal usage ================================//
// ===========================================================================================//

// func generateTrxAddress() {

// }

// func initTrxNetwork() {

// }
