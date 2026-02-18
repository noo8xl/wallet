package config

// // ===========================================================================================//
// // ================================ crypto-api's configs =====================================//
// // ===========================================================================================//

// GetBitcoinAPIKey -> get an BTC api key
func GetBitcoinAPIKey() string {
	return getConfigVar("BTC_API_KEY")
}

// GetEthereumAPIKey -> get an ETH api key
func GetEthereumAPIKey() string {
	return getConfigVar("ETH_API_KEY")
}

// GetLitecoinAPIKey -> get an LTC api key
func GetLitecoinAPIKey() string {
	return getConfigVar("LTC_API_KEY")
}

// GetDogecoinAPIKey -> get an DOGE api key
func GetDogecoinAPIKey() string {
	return getConfigVar("DOGE_API_KEY")
}

// GetTronAPIKey -> get an TRX api key
func GetTronAPIKey() string {
	return getConfigVar("TRX_API_KEY")
}

// GetTheOpenNetworkAPIKey -> get an TON api key
func GetTheOpenNetworkAPIKey() string {
	return ""
}

// GetCryptoAPIKeys -> get all crypto api keys
func GetCryptoAPIKeys() map[string]string {
	return map[string]string{
		"BTC": GetBitcoinAPIKey(),
		"ETH": GetEthereumAPIKey(),
		"TRX": GetTronAPIKey(),
		"TON": GetTheOpenNetworkAPIKey(),
	}
}
