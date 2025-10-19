package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math"
	"math/big"
	"net/http"
	"strings"
	"wallet/pkg/exceptions"
)

// PrintPretty -> print pretty struct variables
func PrintPretty(data any) {
	str, _ := json.MarshalIndent(data, "", "\t")
	fmt.Println(string(str))
}

// ToFixed func -> fix float symbols val after .
func ToFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

// ValidateArgs -> validate range of arguments
func ValidateArgs(argsLen int, expectedValue int) error {
	if argsLen <= 0 || expectedValue <= 0 {
		return errors.New("err: args val our of range")
	}
	if argsLen < expectedValue {
		return errors.New("err: args val our of range")
	}
	return nil
}

func ValidateCard(cardNum string) {
	// https://gocardless.com/guides/posts/what-is-luhn-algorithm/

}

func CheckNetworkConnection() error {
	_, err := http.Get("http://clients3.google.com/generate_204")
	if err != nil {
		return errors.New("network connection failed: " + err.Error())
	}
	return nil
}

// BalanceFromStoreFormatter -> result arg is a string representation of float64 number
func BalanceFromStoreFormatter(result string, e error) *big.Float {

	if e != nil {
		return nil
	}
	if result != "" {
		b := new(big.Float)
		b.SetString(result)
		return b
	}
	return nil
}

// ===========================================================================================//
// ==================================== get coin rates here ==================================//
// ===========================================================================================//

// https://api.coingecko.com/api/v3/coins/markets?vs_currency=${currency}&ids=${coinNameForUrl}
// doc is here -> https://www.coingecko.com/api/documentation <-

// GetRate -> get coin rate in chosen currency by coinName
func GetRate(coinName, currency string) float64 {

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
