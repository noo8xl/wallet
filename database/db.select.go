package database

import (
	"log"
	"strconv"
	"strings"
	"wallet/lib/exceptions"
)

// SelectBtcPrivate -> get private key by address and userId
func (s *DatabaseService) SelectBtcPrivate(address string) string {
	s.db = dbConnect()
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
	s.db = dbConnect()
	var pKey []byte
	defer s.db.Close()

	sql := strings.Join([]string{"SELECT privateKey FROM tonWallets WHERE address=", "\"", address, "\""}, "")
	err := s.db.QueryRow(sql).Scan(&pKey)
	if err != nil {
		panic(err.Error())
	}

	return pKey
}

// IsWalletExists -> check is wallet already exists for current user
// to get a permission to create a permanent wallet for him
func (s *DatabaseService) IsWalletExists(userId int64, bc string) bool {
	s.db = dbConnect()
	var id int64 = 0
	strId := strconv.Itoa(int(userId))
	tableName := s.defineAndGetTableNameByBlockchainShortName(bc)
	defer s.db.Close()

	sql := strings.Join([]string{"SELECT id FROM ", tableName, " WHERE userId=", strId, ";"}, "")
	err := s.db.QueryRow(sql).Scan(&id)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return false
		}
		panic(err.Error())
	}

	return true
}

func (s *DatabaseService) defineAndGetTableNameByBlockchainShortName(bc string) string {

	var tableName string
	switch bc {
	case "btc":
		tableName = "btcWallets"
	case "eth":
		tableName = "ethWallets"
	case "trx":
		tableName = "tronWallets"
	case "ton":
		tableName = "tonWallets"
	case "sol":
		tableName = "solWallets"
	default:
		exceptions.HandleAnException("got a wrong coinname in <getTableNameByBlockchainShortName> func.")
	}

	return tableName
}
