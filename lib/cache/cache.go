package cache

import (
	"log"
	"wallet-cli/config"

	"github.com/go-redis/redis"
)

//
// the doc is -> https://redis.io/docs/latest/operate/oss_and_stack/install/install-redis/
// ------------------------------------------------------------------
// redis with JSON implementation  <-
// https://github.com/redis/go-redis/releases/tag/v9.3.0
//

// SetWalletData -> save new user wallet coin data before saving to db*
func SetWalletData(userId, coinName, address string) {

	c := make(map[string]interface{}, 3)
	rdb := connectRedisStore()
	defer rdb.Close()

	c["UserId"] = userId
	c["CoinName"] = coinName
	c["Address"] = address

	st := rdb.HMSet(userId, c)
	log.Println("redis status -> ", st)

	log.Println("Cache is saved.")
}

// GetWalletData -> get wallet data by userId and coinName
func GetWalletData(userId, coinName string) map[string]string {

	rdb := connectRedisStore()
	defer rdb.Close()

	res := rdb.HGetAll(userId)
	log.Println("cache is -> ", res.Val())
	return res.Val()
}

// ############################################################################################
// ####################### -> connectors to redis stores below: <- ############################
// ############################################################################################

// connectWalletRedisStore -> connect to wallet data storage  func <-
func connectRedisStore() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     config.GetRedisAddr(),
		Password: "",
		DB:       0,
	})
}
