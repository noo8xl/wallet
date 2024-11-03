// Package tron -> is tron network interacting
package tron

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"wallet-cli/config"
	"wallet-cli/lib/exceptions"
	"wallet-cli/lib/models"
)

type trxConfig struct {
	token   string
	network string
}

// https://developers.tron.network/reference/background#note
// -> the doc is here <-
//
// ValidateTrxAddress -> validate address in tron network
func ValidateTrxAddress(addr string) bool {

	cnf := initTrxConfig()
	// fmt.Println("token is => ", cnf.token)

	// testNet + path + payload
	var response struct {
		Result  bool   `json:"result"`
		Message string `json:"message"`
	}
	path := "/wallet/validateaddress"
	payload := strings.NewReader(strings.Join([]string{"{\"address\":", "\"", addr, "\",", "\"visible\":true}"}, ""))
	url := strings.Join([]string{cnf.network, path}, "")

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

func SendSingleTrxTransaction(dto *models.SendTransactionDto) string {

	// save a tsx details to db
	return "hash"
}

// ===========================================================================================//
// ============================== function for internal usage ================================//
// ===========================================================================================//

func initTrxConfig() *trxConfig {
	var conf = new(trxConfig)
	conf.token = config.GetTronAPIKey()
	conf.network = "https://api.trongrid.io" // mainnet
	// conf.netwotk = "https://api.shasta.trongrid.io" // testnet
	//
	return conf
}
