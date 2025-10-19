package redis

import (
	"context"
)

func (s *Store) SetAKey(key, val string) error {

	err := s.rdb.Set(context.Background(), key, val, timeout).Err()
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) GetAKey(key string) (string, error) {

	result, err := s.rdb.Get(context.Background(), key).Result()
	if err != nil {
		return "", err
	}
	return result, nil
}
