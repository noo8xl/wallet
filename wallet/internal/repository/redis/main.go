package redis

import (
	"time"
	"wallet/wallet/config"

	"github.com/redis/go-redis/v9"
)

//
// the doc is -> https://redis.uptrace.dev/guide/go-redis.html

var (
	timeout = time.Millisecond * 3000
)

type Store struct {
	rdb *redis.Client
}

func InitNewStore() *Store {
	rdb := connectRedisStore()
	return &Store{rdb}
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
