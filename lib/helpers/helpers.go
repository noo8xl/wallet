package helpers

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"wallet-cli/lib/exceptions"
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
func ValidateArgs(argsLen int, expectedValue int) {
	if argsLen < expectedValue {
		exceptions.HandleAnException("wrong args range")
	}
}

func ValidateCard(cardNum string) {
	// https://gocardless.com/guides/posts/what-is-luhn-algorithm/

}

func CheckAnInternetConnection() {
	_, err := http.Get("http://clients3.google.com/generate_204")
	if err != nil {
		exceptions.HandleAnException("")
	}
}
