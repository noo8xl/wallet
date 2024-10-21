package cache

//
// the doc is -> https://redis.io/docs/latest/operate/oss_and_stack/install/install-redis/
// ------------------------------------------------------------------
// redis with JSON implementation  <-
// https://github.com/redis/go-redis/releases/tag/v9.3.0
//

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
