package database

import (
	"log"
	"strings"
)

// SelectBtcPrivate -> get private key by address and userId
func (s *DatabaseService) SelectBtcPrivate(address string) string {

	var pKey string
	defer s.db.Close()

	sql := strings.Join([]string{"SELECT privateKey FROM btcWallets WHERE address=", "\"", address, "\""}, "")
	log.Println("sql str -> ", sql)
	err := s.db.QueryRow(sql).Scan(&pKey)
	if err != nil {
		panic(err.Error())
	}

	return pKey
}

// SelectTonPrivate -> get private key by address and userId
func (s *DatabaseService) SelectTonPrivate(address string) []byte {

	var pKey []byte
	defer s.db.Close()

	sql := strings.Join([]string{"SELECT privateKey FROM tonWallets WHERE address=", "\"", address, "\""}, "")
	err := s.db.QueryRow(sql).Scan(&pKey)
	if err != nil {
		panic(err.Error())
	}

	return pKey
}
