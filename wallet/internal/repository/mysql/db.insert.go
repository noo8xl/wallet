package mysql

import (
	"strings"
	models "wallet/wallet/pkg/models/wallet"
)

// doc is -> https://github.com/go-sql-driver/mysql/wiki/Examples

func (s *DatabaseService) InsertBtcWalletToPermanent(dto *models.BtcWallet) error {

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

func (s *DatabaseService) InsertTrxWalletToPermanent(dto *models.TrxWallet) error {

	// * pubKeys are temporary excluded  *
	sql := strings.Join([]string{
		"INSERT INTO trxWallets ",
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

func (s *DatabaseService) InsertTrxWalletToOneTimeAddresses(dto *models.TrxWallet) error {

	// * pubKeys are temporary excluded  *
	sql := strings.Join([]string{
		"INSERT INTO oneTimeTrxWallets ",
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

func (s *DatabaseService) InsertLtcWalletToPermanent(dto *models.LtcWallet) error {

	// * pubKeys are temporary excluded  *
	sql := strings.Join([]string{
		"INSERT INTO ltcWallets ",
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

func (s *DatabaseService) InsertLtcWalletToOneTimeAddresses(dto *models.LtcWallet) error {
	// * pubKeys are temporary excluded  *
	sql := strings.Join([]string{
		"INSERT INTO oneTimeLtcWallets ",
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

func (s *DatabaseService) InsertDogeWalletToPermanent(dto *models.Dogewallet) error {
	// * pubKeys are temporary excluded  *
	sql := strings.Join([]string{
		"INSERT INTO dogeWallets ",
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

func (s *DatabaseService) InsertDogeWalletToOneTimeAddresses(dto *models.Dogewallet) error {

	// * pubKeys are temporary excluded  *
	sql := strings.Join([]string{
		"INSERT INTO oneTimeDogeWallets ",
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

// InsertTrxWallet -> insert user wallet data to db
// func InsertTrxWallet(dto models.TonWallet) error {

// 	// ctx := context.Background()

// 	// should be updated
// 	sql := "INSERT INTO tonWallets (address, addrType, privateKey, bitsLen, userId) VALUES (?,?,?,?,?)"
// 	res := db.QueryRow(sql, dto.Address, dto.AddrType, dto.PrivateKey, dto.BitsLen, dto.UserId)
// fmt.Println("sql result is -> ", res)

// return res.Err()
// }
