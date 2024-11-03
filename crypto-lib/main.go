package cryptolib

import (
	"strings"
	"sync"
	"wallet-cli/api"
	"wallet-cli/config"
	"wallet-cli/crypto-lib/bitcoin"
	"wallet-cli/crypto-lib/ethereum"
	theopennetwork "wallet-cli/crypto-lib/the-open-network"
	"wallet-cli/crypto-lib/tron"
	"wallet-cli/lib/exceptions"
	"wallet-cli/lib/models"
)

// CreateWalletList -> create a pool of the workers and
// create a wallet for the user in each available blockchain
func CreateWalletList(userId *string) []*models.WalletListItem {

	result := make(chan *models.WalletListItem, 4)
	var walletItem *models.WalletListItem
	var wg sync.WaitGroup
	coinList := config.GetAvailableCoinList()
	var walletList []*models.WalletListItem

	for _, item := range coinList {
		wg.Add(1)
		opt := &routineOpts{
			userId:   *userId,
			coinName: item,
			result:   result,
			wg:       &wg,
		}
		go worker(opt)
	}

	wg.Wait()
	close(result)

	for i := 0; i < len(coinList); i++ {
		walletItem = <-result
		if walletItem != nil {
			walletList = append(walletList, walletItem)
		}
	}

	return walletList
}

// DefineAndRunBlockchain -> define a blockchain, init connection
// and then generate a new address for the user
func DefineAndRunBlockchain(coin, userId *string) *models.WalletListItem {

	var walletItem *models.WalletListItem

	switch strings.ToLower(*coin) {
	case "btc":
		walletItem = bitcoin.CreateWallet(userId)
	case "eth":
		// walletItem = ethereum.CreateWallet(userId)
		walletItem = &models.WalletListItem{CoinName: "eth", Address: "will be"}
	case "trx":
		walletItem = tron.CreateWallet(userId)
	case "ton":
		walletItem = theopennetwork.CreateWallet(userId)
	default:
		exceptions.HandleAnException("Got an unknown blockchain in <get wallet>")
	}

	return walletItem
}

// DefineBlockchainAndGetBalance -> define a blockchain, init connection
// and get a balance in crypto by address and then get a fiat balance
// by received crypto amount in chosen currencyType
func DefineBlockchainAndGetBalance(coin, address, currencyType string) *models.ResponseBalance {

	var fiatBalance float64
	var response = new(models.ResponseBalance)
	response.CoinName = coin
	response.CurrencyType = currencyType

	switch coin {
	case "btc":
		response.CoinBalance = bitcoin.GetBitcoinAddressBalance(address)
		fiatBalance = api.GetRate("bitcoin", currencyType)
	case "ton":
		response.CoinBalance = theopennetwork.GetTonBalanceByAddress(address)
		fiatBalance = api.GetRate("the-open-network", currencyType)
	case "eth":
		response.CoinBalance = ethereum.GetEthBalanceByAddress(address)
		fiatBalance = api.GetRate("ethereum", currencyType)
	case "trx":
		response.CoinBalance = theopennetwork.GetTonBalanceByAddress(address)
		fiatBalance = api.GetRate("tron", currencyType)
	default:
		exceptions.HandleAnException("Got an unknown blockchain at the <get wallet balance>")
	}

	response.FiatAmount = getaBigFloat(fiatBalance, response.CoinBalance)
	return response
}

// DefineBlockchainAndSendTsx -> define a blockchain, init connection
// and send transaction to user by address and coinName
func DefineBlockchainAndSendTsx(dto *models.SendTransactionDto) string {

	var hash string

	switch dto.CoinName {
	case "btc":
		hash = bitcoin.SendSingleBtcTransaction(dto)
	case "eth":
		hash = ethereum.SendSingleEthTransaction(dto)
	case "ton":
		hash = theopennetwork.SendSingleTonTransaction(dto)
	case "trx":
		hash = tron.SendSingleTrxTransaction(dto)
	default:
		exceptions.HandleAnException("Got an unknown blockchain at the <send transaction>")
	}

	return hash
}
