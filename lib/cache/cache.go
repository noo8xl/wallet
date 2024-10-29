package cache

//
// the doc is -> https://redis.io/docs/latest/operate/oss_and_stack/install/install-redis/

// ############################################################################################
// ####################### -> connectors to redis stores below: <- ############################
// ############################################################################################

// connectWalletRedisStore -> connect to wallet data storage  func <-
// func connectRedisStore() *redis.Client {
// 	return redis.NewClient(&redis.Options{
// 		Addr:     config.GetRedisAddr(),
// 		Password: "",
// 		DB:       0,
// 	})
// }
