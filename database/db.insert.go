package database

import (
	"strings"
	"wallet/lib/models"
)

// doc is -> https://github.com/go-sql-driver/mysql/wiki/Examples

func (s *DatabaseService) InsertBtcWalletToPermanent(dto *models.BtcWallet) error {
	s.db = dbConnect()
	defer s.db.Close()

	// * pubKeys are temporary excluded  *
	sql := strings.Join([]string{
		"INSERT INTO btcWallets ",
		"(address, privateKey, publicKey, wif, scriptType, originalAddress, oapAddress, userId) ",
		"VALUES (?,?,?,?,?,?,?,?)",
	}, "")
	res := s.db.QueryRow(
		sql, dto.Address, dto.PrivateKey,
		dto.PublicKey, dto.Wif, dto.ScriptType,
		dto.OriginalAddress, dto.OAPAddress, dto.CustomerId,
	)
	// fmt.Println("sql result is -> ", &res)

	return res.Err()
}

func (s *DatabaseService) InsertBtcWalletToOneTimeAddresses(dto *models.BtcWallet) error {
	s.db = dbConnect()
	defer s.db.Close()

	// * pubKeys are temporary excluded  *
	sql := strings.Join([]string{
		"INSERT INTO oneTimeBtcWallets ",
		"(address, privateKey, publicKey, wif, userId) ",
		"VALUES (?,?,?,?,?)",
	}, "")
	res := s.db.QueryRow(
		sql, dto.Address, dto.PrivateKey,
		dto.PublicKey, dto.Wif, dto.CustomerId,
	)
	// fmt.Println("sql result is -> ", &res)

	return res.Err()
}

// InsertEthWallet -> insert user wallet data to db
func (s *DatabaseService) InsertEthWalletToPermament(dto *models.EthWallet) error {
	s.db = dbConnect()
	defer s.db.Close()

	// * pubKeys are temporary excluded  *
	sql := strings.Join([]string{
		"INSERT INTO ethWallets ",
		"(address, privateKey, publicKey, wif, scriptType, originalAddress, oapAddress, userId) ",
		"VALUES (?,?,?,?,?,?,?,?)",
	}, "")
	res := s.db.QueryRow(
		sql, dto.Address, dto.PrivateKey,
		dto.PublicKey, dto.Wif, dto.ScriptType,
		dto.OriginalAddress, dto.OAPAddress, dto.CustomerId,
	)

	// fmt.Println("sql result is -> ", &res)

	return res.Err()
}

func (s *DatabaseService) InsertEthWalletToOneTimeAddresses(dto *models.EthWallet) error {
	s.db = dbConnect()
	defer s.db.Close()

	// * pubKeys are temporary excluded  *
	sql := strings.Join([]string{
		"INSERT INTO oneTimeEthWallets ",
		"(address, privateKey, publicKey, wif, userId) ",
		"VALUES (?,?,?,?,?)",
	}, "")
	res := s.db.QueryRow(
		sql, dto.Address, dto.PrivateKey,
		dto.PublicKey, dto.Wif, dto.CustomerId,
	)
	// fmt.Println("sql result is -> ", &res)

	return res.Err()
}

// InsertTonWallet -> insert user wallet data to db
func (s *DatabaseService) InsertTonWalletToPermanent(dto *models.TonWallet) error {
	s.db = dbConnect()
	// ctx := context.Background()
	defer s.db.Close()

	// should be updated
	sql := strings.Join([]string{
		"INSERT INTO tonWallets ",
		"(address, addrType, privateKey, bitsLen, userId) ",
		"VALUES (?,?,?,?,?)",
	}, "")
	res := s.db.QueryRow(
		sql, dto.Address, dto.AddrType,
		dto.PrivateKey, dto.BitsLen, dto.CustomerId,
	)
	// fmt.Println("sql result is -> ", &res)

	return res.Err()
}

func (s *DatabaseService) InsertTonWalletToOneTimeAddresses(dto *models.TonWallet) error {
	s.db = dbConnect()
	// ctx := context.Background()
	defer s.db.Close()

	// should be updated
	sql := strings.Join([]string{
		"INSERT INTO oneTimeTonWallets ",
		"(address, addrType, privateKey, bitsLen, userId) ",
		"VALUES (?,?,?,?,?)",
	}, "")
	res := s.db.QueryRow(
		sql, dto.Address, dto.AddrType,
		dto.PrivateKey, dto.BitsLen, dto.CustomerId,
	)
	// fmt.Println("sql result is -> ", &res)

	return res.Err()
}

// InsertTrxWallet -> insert user wallet data to db
// func InsertTrxWallet(dto models.TonWallet) error {

// 	// ctx := context.Background()
// 	db := dbConnect()
// 	defer db.Close()

// 	// should be updated
// 	sql := "INSERT INTO tonWallets (address, addrType, privateKey, bitsLen, userId) VALUES (?,?,?,?,?)"
// 	res := db.QueryRow(sql, dto.Address, dto.AddrType, dto.PrivateKey, dto.BitsLen, dto.UserId)
// fmt.Println("sql result is -> ", res)

// return res.Err()
// }
