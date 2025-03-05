package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	timeout = time.Millisecond * 3000
)

type Store struct {
	rdb *redis.Client
}

func InitNewStore() *Store {
	return &Store{
		rdb: &redis.Client{},
	}
}

func (s *Store) SetAKey(key, val string) error {
	s.rdb = connectRedisStore()
	defer s.rdb.Close()
	err := s.rdb.Set(context.Background(), key, val, timeout).Err()
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) GetAKey(key string) (string, error) {
	s.rdb = connectRedisStore()
	defer s.rdb.Close()
	result, err := s.rdb.Get(context.Background(), key).Result()
	if err != nil {
		return "", err
	}
	return result, nil
}
