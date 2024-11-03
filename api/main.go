package api

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"wallet-cli/lib/exceptions"
)

// ===========================================================================================//
// ==================================== get coin rates here ==================================//
// ===========================================================================================//

// https://api.coingecko.com/api/v3/coins/markets?vs_currency=${currency}&ids=${coinNameForUrl}
// doc is here -> https://www.coingecko.com/api/documentation <-

// GetRate -> get coin rate in chosen currency by coinName
func GetRate(coinName string, currency string) float64 {

	var resp map[string]map[string]float64
	uri := strings.Join([]string{"https://api.coingecko.com/api/v3/simple/price?ids=", coinName, "&vs_currencies=", currency}, "")
	// fmt.Println("cur uri is ->\n", uri)

	res, err := http.Get(uri)
	if err != nil {
		exceptions.HandleAnException("<GetRate> got an http err: " + err.Error())
	}
	defer res.Body.Close()

	resBody, _ := io.ReadAll(res.Body)
	json.Unmarshal(resBody, &resp)
	// log.Println("resp  -> ", resp[coinName][currency])

	return resp[coinName][currency]
}
