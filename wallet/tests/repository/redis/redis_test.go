package cache

import (
	"testing"
	redis "wallet/wallet/internal/repository/redis"
)

var (
	store redis.Store
)

func TestInitRedisStore(t *testing.T) {
	store := redis.InitNewStore()
	if store == nil {
		t.Fatalf("init rdb connection failed")
	}

	t.Logf("store is -> %v", store)
}

func TestSetAKey(t *testing.T) {
	store := redis.InitNewStore()
	err := store.SetAKey("1", "test_chached_value")
	if err != nil {
		t.Fatalf("TestSetAKey failed with -> %v", err.Error())
	}
}

func TestGetAKey(t *testing.T) {
	store := redis.InitNewStore()
	result, err := store.GetAKey("1")
	if err != nil {
		t.Fatalf("TestSetAKey failed with -> %v", err.Error())
	}

	t.Logf("received key is -> %s", result)
}
