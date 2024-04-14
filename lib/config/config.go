package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

// GetSQLDatabaseConfig -> get mysql db config from env file
func GetSQLDatabaseConfig() string {

	user := getConfigVar("DB_USER")
	pwd := getConfigVar("DB_PASSWORD")
	name := getConfigVar("DB_NAME")

	return strings.Join([]string{user, ":", pwd, "@/", name}, "")
}

// GetAvailableCoinList -> get list of available coins
func GetAvailableCoinList() [2]string {
	return [2]string{
		"BTC", "TON",
		// "ETH", "TRX",
	}
}

// GetRedisAddr -> return a redis config str*
func GetRedisAddr() string {
	return getConfigVar("REDIS_ADDR")
}

// ########################################################################################### //
// ################################ internal usage only ###################################### //
// ########################################################################################### //

func getConfigVar(key string) string {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println(err)
	}
	return os.Getenv(key)
}
