package mysql

import (
	"errors"
	"strconv"
	"strings"
	"wallet/wallet/config"
	"wallet/wallet/pkg/utils"
)

func (s *DatabaseService) DefineaBlockchainAndGetPrivateKey(address string, bc string, opts byte) (string, error) {
	var err error
	var privateKey string
	var encryptedPk string

	key := config.GetAnEncryptionKey()

	switch bc {
	case "btc":
		encryptedPk, err = s.selectBtcPrivate(address, opts)
	case "eth":
		encryptedPk, err = s.selectEthPrivate(address, opts)
	case "ltc":
		encryptedPk, err = s.selectLtcPrivate(address, opts)
	case "doge":
		encryptedPk, err = s.selectDogePrivate(address, opts)
	case "trx":
		encryptedPk, err = s.selectTrxPrivate(address, opts)
	case "ton":
		encryptedPk, err = s.selectTonPrivate(address, opts)
	case "sol":
		encryptedPk, err = s.selectSolPrivate(address, opts)
	default:
		return "", errors.New("got a wrong coinname")
	}

	if err != nil {
		return "", err
	}

	privateKey, err = utils.DecryptKey(key, encryptedPk)
	if err != nil {
		return "", err
	}

	return privateKey, nil

}

// IsWalletExists -> check is wallet already exists for current user
// to get a permission to create a permanent wallet for him
func (s *DatabaseService) IsWalletExists(userId int64, bc string, opts byte) (bool, error) {

	var id int64 = 0
	var err error
	var tableName string = ""
	strId := strconv.Itoa(int(userId))

	switch opts {
	case 0:
		tableName, err = s.defineAndGetPermanentTableNameByBlockchainShortName(bc)
		if err != nil {
			return false, err
		}
	case 1:
		tableName, err = s.defineAndGetOneTimeTableNameByBlockchainShortName(bc)
		if err != nil {
			return false, err
		}
	default:
		return false, errors.New("got a wrong option value")
	}

	sql := strings.Join([]string{"SELECT id FROM ", tableName, " WHERE userId=\"", strId, "\";"}, "")
	err = s.db.QueryRow(sql).Scan(&id)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

// ######################################################################################################
// ######################################################################################################
// ######################################################################################################

// SelectBtcPrivate -> get private key by address and userId
func (s *DatabaseService) selectBtcPrivate(address string, opts byte) (string, error) {

	var tableName string
	var pKey string

	switch opts {
	case 0:
		tableName = "btcWallets"
	case 1:
		tableName = "oneTimeBtcWallets"
	default:
		return "", errors.New("got a wrong option value")
	}

	sql := strings.Join([]string{"SELECT privateKey FROM ", tableName, " WHERE address=", "\"", address, "\""}, "")
	err := s.db.QueryRow(sql).Scan(&pKey)
	if err != nil {
		return "", err
	}

	return pKey, nil
}

// SelectBtcPrivate -> get private key by address and userId
func (s *DatabaseService) selectEthPrivate(address string, opts byte) (string, error) {

	var tableName string
	var pKey string

	switch opts {
	case 0:
		tableName = "ethWallets"
	case 1:
		tableName = "oneTimeEthWallets"
	default:
		return "", errors.New("got a wrong option value")
	}

	sql := strings.Join([]string{"SELECT privateKey FROM ", tableName, " WHERE address=", "\"", address, "\""}, "")
	err := s.db.QueryRow(sql).Scan(&pKey)
	if err != nil {
		return "", err
	}

	return pKey, nil
}

// SelectTrxPrivate -> get private key by address and userId
func (s *DatabaseService) selectTrxPrivate(address string, opts byte) (string, error) {

	var tableName string
	var pKey string

	switch opts {
	case 0:
		tableName = "trxWallets"
	case 1:
		tableName = "oneTimeTrxWallets"
	default:
		return "", errors.New("got a wrong option value")
	}

	sql := strings.Join([]string{"SELECT privateKey FROM ", tableName, " WHERE address=", "\"", address, "\""}, "")
	err := s.db.QueryRow(sql).Scan(&pKey)
	if err != nil {
		return "", err
	}

	return pKey, nil
}

// SelectTonPrivate -> get private key by address and userId
func (s *DatabaseService) selectTonPrivate(address string, opts byte) (string, error) {

	var tableName string
	var pKey []byte

	switch opts {
	case 0:
		tableName = "tonWallets"
	case 1:
		tableName = "oneTimeTonWallets"
	default:
		return "", errors.New("got a wrong option value")
	}

	sql := strings.Join([]string{"SELECT privateKey FROM ", tableName, " WHERE address=", "\"", address, "\""}, "")
	err := s.db.QueryRow(sql).Scan(&pKey)
	if err != nil {
		return "", err
	}

	return string(pKey), nil
}

func (s *DatabaseService) selectLtcPrivate(address string, opts byte) (string, error) {

	var pKey string
	var tableName string

	switch opts {
	case 0:
		tableName = "ltcWallets"
	case 1:
		tableName = "oneTimeLtcWallets"
	default:
		return "", errors.New("got a wrong option value")
	}

	sql := strings.Join([]string{"SELECT privateKey FROM ", tableName, " WHERE address=", "\"", address, "\""}, "")
	err := s.db.QueryRow(sql).Scan(&pKey)
	if err != nil {
		return "", err
	}

	return pKey, nil
}

func (s *DatabaseService) selectDogePrivate(address string, opts byte) (string, error) {

	var tableName string
	var pKey string

	switch opts {
	case 0:
		tableName = "dogeWallets"
	case 1:
		tableName = "oneTimeDogeWallets"
	default:
		return "", errors.New("got a wrong option value")
	}

	sql := strings.Join([]string{"SELECT privateKey FROM ", tableName, " WHERE address=", "\"", address, "\""}, "")
	err := s.db.QueryRow(sql).Scan(&pKey)
	if err != nil {
		return "", err
	}

	return pKey, nil
}

func (s *DatabaseService) selectSolPrivate(address string, opts byte) (string, error) {

	var tableName string
	var pKey string

	switch opts {
	case 0:
		tableName = "solWallets"
	case 1:
		tableName = "oneTimeSolWallets"
	default:
		return "", errors.New("got a wrong option value")
	}

	sql := strings.Join([]string{"SELECT privateKey FROM ", tableName, " WHERE address=", "\"", address, "\""}, "")
	err := s.db.QueryRow(sql).Scan(&pKey)
	if err != nil {
		return "", err
	}

	return pKey, nil
}

func (s *DatabaseService) defineAndGetPermanentTableNameByBlockchainShortName(bc string) (string, error) {

	var tableName string
	switch bc {
	case "btc":
		tableName = "btcWallets"
	case "eth":
		tableName = "ethWallets"
	case "ltc":
		tableName = "ltcWallets"
	case "doge":
		tableName = "dogeWallets"
	case "trx":
		tableName = "tronWallets"
	case "ton":
		tableName = "tonWallets"
	case "sol":
		tableName = "solWallets"
	default:
		return "", errors.New("got a wrong coinname in <defineAndGetPermanentTableNameByBlockchainShortName>")
	}

	return tableName, nil
}

func (s *DatabaseService) defineAndGetOneTimeTableNameByBlockchainShortName(bc string) (string, error) {

	var tableName string
	switch bc {
	case "btc":
		tableName = "oneTimeBtcWallets"
	case "eth":
		tableName = "oneTimeEthWallets"
	case "ltc":
		tableName = "oneTimeLtcWallets"
	case "doge":
		tableName = "oneTimeDogeWallets"
	case "trx":
		tableName = "oneTimeTronWallets"
	case "ton":
		tableName = "oneTimeTonWallets"
	case "sol":
		tableName = "oneTimeSolWallets"
	default:
		return "", errors.New("got a wrong coinname in <defineAndGetOneTimeTableNameByBlockchainShortName>")
	}

	return tableName, nil
}
