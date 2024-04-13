package api

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strings"
)

// ===========================================================================================//
// ==================================== get coin rates here ==================================//
// ===========================================================================================//

// https://api.coingecko.com/api/v3/coins/markets?vs_currency=${currency}&ids=${coinNameForUrl}

// doc is here -> https://www.coingecko.com/api/documentation <-

func GetRate(coinName, currency string) float64 {

	var rate float64
	var resp []map[string]interface{}
	url := strings.Join([]string{"https://api.coingecko.com/api/v3/coins/markets?vs_currency=", currency, "&ids=", coinName}, "")
	// fmt.Println("reseived data is ->\n", url)

	res, err := http.Get(url)
	if err != nil {
		os.Exit(1)
	}
	defer res.Body.Close()

	resBody, _ := io.ReadAll(res.Body)
	json.Unmarshal(resBody, &resp)

	rate = resp[0]["current_price"].(float64)
	// log.Println("rate -> ", rate)

	return rate
}
