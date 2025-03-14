package common

import (
	"math/big"
	"strings"
	"sync"
	pb "wallet/api"
	"wallet/config"
	"wallet/crypto-lib/bitcoin"
	"wallet/crypto-lib/dogecoin"
	"wallet/crypto-lib/ethereum"
	"wallet/crypto-lib/litecoin"
	"wallet/crypto-lib/solana"
	theopennetwork "wallet/crypto-lib/the-open-network"
	"wallet/crypto-lib/tron"
	"wallet/lib/exceptions"
)

type Service struct {
	btcService  *bitcoin.Service
	ltcService  *litecoin.Service
	dogeService *dogecoin.Service
	ethService  *ethereum.Service
	tonService  *theopennetwork.Service
	solService  *solana.Service
	trxService  *tron.Service
}

func InitService() *Service {
	return &Service{
		btcService:  bitcoin.InitService(),
		ltcService:  litecoin.InitService(),
		dogeService: dogecoin.InitService(),
		ethService:  ethereum.InitService(),
		trxService:  tron.InitService(),
		solService:  solana.InitService(),
		tonService:  theopennetwork.InitService(),
	}
}

// CreateWalletList -> create a pool of the workers and
// create a wallet for the user in each available blockchain
func (s *Service) CreateWalletList(userId int64) *pb.WalletList {

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

	for range len(coinList) {
		walletItem = <-result
		if walletItem != nil {
			walletList.Wallet = append(walletList.Wallet, walletItem)
		}
	}

	return walletList
}

// DefineAndRunBlockchain -> define a blockchain, init connection
// and then generate a new address for the user
func (s *Service) DefineBlockchainAndCreatePermanentWallet(coin string, userId int64) *pb.WalletItem {

	var walletItem *pb.WalletItem

	switch strings.ToLower(coin) {
	case "btc":
		walletItem = s.btcService.CreatePermanentWallet(userId)
	case "ltc":
		walletItem = s.ltcService.CreatePermanentWallet(userId)
	case "doge":
		walletItem = s.dogeService.CreatePermanentWallet(userId)
	case "eth":
		walletItem = s.ethService.CreatePermanentWallet(userId)
	case "trx":
		walletItem = s.trxService.CreatePermanentWallet(userId)
	case "sol":
		walletItem = s.solService.CreatePermanentWallet(userId)
	case "ton":
		walletItem = s.tonService.CreatePermanentWallet(userId)
	default:
		exceptions.HandleAnException("Got an unknown blockchain in <get wallet>")
	}

	return walletItem
}

// DefineAndRunBlockchain -> define a blockchain, init connection
// and then generate a new address for the user
func (s *Service) DefineBlockchainAndCreateOneTimeAddress(coin string, userId int64) *pb.WalletItem {

	var walletItem *pb.WalletItem

	switch strings.ToLower(coin) {
	case "btc":
		walletItem = s.btcService.CreateOneTimeddress(userId)
	case "ltc":
		walletItem = s.ltcService.CreateOneTimeddress(userId)
	case "doge":
		walletItem = s.dogeService.CreateOneTimeddress(userId)
	case "eth":
		walletItem = s.ethService.CreateOneTimeddress(userId)
	case "trx":
		walletItem = s.trxService.CreateOneTimeddress(userId)
	case "ton":
		walletItem = s.tonService.CreateOneTimeddress(userId)
	default:
		exceptions.HandleAnException("Got an unknown blockchain in <get wallet>")
	}

	return walletItem
}

// DefineBlockchainAndGetCoinBalance -> define a blockchain, init connection
// and get a balance in crypto by address and coin name
func (s *Service) DefineBlockchainAndGetCoinBalance(coin, address string) *pb.CoinBalance {

	var balance string

	switch coin {
	case "btc":
		balance = s.btcService.GetBalanceByAddress(address).String()
	case "ltc":
		balance = s.ltcService.GetBalanceByAddress(address).String()
	case "doge":
		balance = s.dogeService.GetBalanceByAddress(address).String()
	case "ton":
		balance = s.tonService.GetBalanceByAddress(address).String()
	case "eth":
		balance = s.ethService.GetBalanceByAddress(address).String()
	case "trx":
		balance = s.trxService.GetBalanceByAddress(address).String()
	case "sol":
		balance = s.solService.GetBalanceByAddress(address) // as *big.Float
	default:
		exceptions.HandleAnException("Got an unknown blockchain at the <get wallet balance>")
	}

	return &pb.CoinBalance{CoinName: coin, Balance: balance}
}

func (s *Service) DefineBlockchainAndGetCustomerBalance(currencyType string, customerId int64) *pb.CustomerBalance {

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
func (s *Service) DefineBlockchainAndSendSingleTsx(dto *pb.SendSingleTsxRequest) *pb.TransactionHash {

	var hash string

	switch dto.CoinName {
	case "btc":
		hash = s.btcService.SendSingleTransaction(dto)
	case "ltc":
		hash = s.ltcService.SendSingleTransaction(dto)
	case "doge":
		hash = s.dogeService.SendSingleTransaction(dto)
	case "eth":
		hash = s.ethService.SendSingleTransaction(dto)
	case "ton":
		hash = s.tonService.SendSingleTransaction(dto)
	case "trx":
		hash = s.trxService.SendSingleTransaction(dto)
	case "sol":
		hash = s.solService.SendSingleTransaction(dto)
	default:
		exceptions.HandleAnException("Got an unknown blockchain at the <send transaction>")
	}

	return &pb.TransactionHash{TsxHash: hash}
}

// DefineBlockchainAndSendMultiptleTsx -> define a blockchain, init connection
// and send transaction to the list of users by address and coinName
func (s *Service) DefineBlockchainAndSendMultiptleTsx(dto *pb.SendMultipleTsxRequest) *pb.TransactionHash {

	// TODO: impl this method

	var hash string

	switch dto.CoinName {
	case "btc":
		// hash = s.btcService.SentMu(dto)
	case "ltc":
		//
	case "doge":
		//
	case "eth":
		// hash = s.EthereumService.SendSingleEthTransaction(dto)
	case "ton":
		// hash = s.TONService.SendSingleTonTransaction(dto)
	case "trx":
	// hash = s.TronService.SendSingleTrxTransaction(dto)
	case "sol":
		// hash = s.SolanaService.SendSingleSolTransaction(dto)"
	default:
		exceptions.HandleAnException("Got an unknown blockchain at the <send transaction>")
	}

	return &pb.TransactionHash{TsxHash: hash}
}
