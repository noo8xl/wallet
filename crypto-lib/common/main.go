package common

import (
	"math/big"
	"strings"
	"sync"
	pb "wallet/api"
	"wallet/config"
	"wallet/crypto-lib/bitcoin"
	"wallet/crypto-lib/ethereum"
	theopennetwork "wallet/crypto-lib/the-open-network"
	"wallet/crypto-lib/tron"
	"wallet/lib/exceptions"
)

type ServiceType struct {
	bitcoin.BitcoinService
	ethereum.EthereumService
	theopennetwork.TONService
	tron.TronService
}

// CreateWalletList -> create a pool of the workers and
// create a wallet for the user in each available blockchain
func (s *ServiceType) CreateWalletList(userId int64) *pb.WalletList {

	result := make(chan *pb.WalletItem, 4)
	var walletItem *pb.WalletItem
	var wg sync.WaitGroup
	coinList := config.GetAvailableCoinList()
	walletList := &pb.WalletList{
		Wallet: make([]*pb.WalletItem, 4),
	}

	for _, item := range coinList {
		wg.Add(1)
		opt := &routineOpts{
			userId:   userId,
			coinName: item,
			result:   result,
			wg:       &wg,
		}
		go s.worker(opt)
	}

	wg.Wait()
	close(result)

	for i := 0; i < len(coinList); i++ {
		walletItem = <-result
		if walletItem != nil {
			walletList.Wallet = append(walletList.Wallet, walletItem)
		}
	}

	return walletList
}

// DefineAndRunBlockchain -> define a blockchain, init connection
// and then generate a new address for the user
func (s *ServiceType) DefineBlockchainAndCreateWallet(coin string, userId int64) *pb.WalletItem {

	var walletItem *pb.WalletItem

	switch strings.ToLower(coin) {
	case "btc":
		walletItem = s.BitcoinService.CreateWallet(userId)
		// walletItem = &models.WalletListItem{CoinName: "btc", Address: "addressBTCHere123123_test"}
	case "eth":
		walletItem = s.EthereumService.CreateWallet(userId)
		// walletItem = &models.WalletListItem{CoinName: "eth", Address: "addressETHHere568568_test"}
	case "trx":
		walletItem = s.TronService.CreateWallet(userId)
		// walletItem = &models.WalletListItem{CoinName: "trx", Address: "addressTRXHere1249728_test"}
	case "ton":
		walletItem = s.TONService.CreateWallet(userId)
		// walletItem = &models.WalletListItem{CoinName: "ton", Address: "addressTONHere56-0sj8_test"}
	default:
		exceptions.HandleAnException("Got an unknown blockchain in <get wallet>")
	}

	return walletItem
}

// DefineBlockchainAndGetCoinBalance -> define a blockchain, init connection
// and get a balance in crypto by address and coin name
func (s *ServiceType) DefineBlockchainAndGetCoinBalance(coin, address string) *pb.CoinBalance {

	var balance string

	switch coin {
	case "btc":
		balance = s.BitcoinService.GetBitcoinAddressBalance(address).String()
	case "ton":
		balance = s.TONService.GetTonBalanceByAddress(address).String()
	case "eth":
		balance = s.EthereumService.GetEthBalanceByAddress(address).String()
	case "trx":
		// balance = s.TronService.GetTonBalanceByAddress(address).String()
	default:
		exceptions.HandleAnException("Got an unknown blockchain at the <get wallet balance>")
	}

	return &pb.CoinBalance{CoinName: coin, Balance: balance}
}

func (s *ServiceType) DefineBlockchainAndGetCustomerBalance(currencyType string, customerId int64) *pb.CustomerBalance {

	// TODO: get value from the db by customer

	var fiatBalance float64
	var coinBalance *big.Float

	// coinBalance += each blockchain balance
	// fiatBalance get from coinBalance

	// exceptions.HandleAnException("Got an unknown blockchain at the <get wallet balance>")

	fiat := getaBigFloat(fiatBalance, coinBalance).String()

	return &pb.CustomerBalance{
		CurrencyName: currencyType,
		CoinBalance:  coinBalance.String(),
		FiatBalance:  fiat,
	}
}

// DefineBlockchainAndSendTsx -> define a blockchain, init connection
// and send transaction to user by address and coinName
func (s *ServiceType) DefineBlockchainAndSendSingleTsx(dto *pb.SendSingleTsxRequest) *pb.TransactionHash {

	var hash string

	switch dto.CoinName {
	case "btc":
		hash = s.BitcoinService.SendSingleBtcTransaction(dto)
	case "eth":
		hash = s.EthereumService.SendSingleEthTransaction(dto)
	case "ton":
		hash = s.TONService.SendSingleTonTransaction(dto)
	case "trx":
		hash = s.TronService.SendSingleTrxTransaction(dto)
	default:
		exceptions.HandleAnException("Got an unknown blockchain at the <send transaction>")
	}

	return &pb.TransactionHash{TsxHash: hash}
}

// DefineBlockchainAndSendMultiptlyTsx -> define a blockchain, init connection
// and send transaction to the list of users by address and coinName
func (s *ServiceType) DefineBlockchainAndSendMultiptlyTsx(dto *pb.SendMultiplyTsxRequest) *pb.TransactionHash {

	// TODO: impl this method

	var hash string

	switch dto.CoinName {
	case "btc":
		// hash = s.BitcoinService.SendSingleBtcTransaction(dto)
	case "eth":
		// hash = s.EthereumService.SendSingleEthTransaction(dto)
	case "ton":
		// hash = s.TONService.SendSingleTonTransaction(dto)
	case "trx":
		// hash = s.TronService.SendSingleTrxTransaction(dto)
	default:
		exceptions.HandleAnException("Got an unknown blockchain at the <send transaction>")
	}

	return &pb.TransactionHash{TsxHash: hash}
}
