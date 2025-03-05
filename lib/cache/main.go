package cache

import (
	"wallet/config"

	"github.com/redis/go-redis/v9"
)

//
// the doc is -> https://redis.uptrace.dev/guide/go-redis.html

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
