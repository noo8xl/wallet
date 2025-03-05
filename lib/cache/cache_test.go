package cache

import "testing"

func TestInitRedisStore(t *testing.T) {
	store := InitNewStore()
	if store.rdb == nil {
		t.Fatalf("init rdb connection failed")
	}

	t.Logf("store is -> %v", store.rdb)
}

func TestSetAKey(t *testing.T) {
	store := InitNewStore()
	err := store.SetAKey("1", "test_chached_value")
	if err != nil {
		t.Fatalf("TestSetAKey failed with -> %v", err.Error())
	}
}

func TestGetAKey(t *testing.T) {
	store := InitNewStore()
	result, err := store.GetAKey("1")
	if err != nil {
		t.Fatalf("TestSetAKey failed with -> %v", err.Error())
	}

	t.Logf("received key is -> %s", result)
}
