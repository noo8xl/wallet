package crypto

import (
	"errors"
	"log"
	"math/big"
	"strings"
	"sync"

	pb "wallet/gen/wallet"
	"wallet/wallet/config"
	"wallet/wallet/internal/services/crypto/networks/bitcoin"
	"wallet/wallet/internal/services/crypto/networks/dogecoin"
	"wallet/wallet/internal/services/crypto/networks/ethereum"
	"wallet/wallet/internal/services/crypto/networks/litecoin"
	"wallet/wallet/internal/services/crypto/networks/solana"
	"wallet/wallet/internal/services/crypto/networks/theopennetwork"
	"wallet/wallet/internal/services/crypto/networks/tron"
)

type CryptoService struct {
	btcService  *bitcoin.Service
	ltcService  *litecoin.Service
	dogeService *dogecoin.Service
	ethService  *ethereum.Service
	tonService  *theopennetwork.Service
	solService  *solana.Service
	trxService  *tron.Service
}

func InitService() *CryptoService {
	return &CryptoService{
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
func (s *CryptoService) CreateWalletList(userId int64) *pb.WalletList {

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
func (s *CryptoService) DefineBlockchainAndCreatePermanentWallet(coin string, userId int64) *pb.WalletItem {

	var walletItem *pb.WalletItem
	var err error

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
		err = errors.New("invalid coin name")
	}

	if err != nil {
		log.Panic(err)
		return nil
	}

	return walletItem
}

// DefineAndRunBlockchain -> define a blockchain, init connection
// and then generate a new address for the user
func (s *CryptoService) DefineBlockchainAndCreateOneTimeAddress(coin string, userId int64) *pb.WalletItem {

	var walletItem *pb.WalletItem
	var err error

	switch strings.ToLower(coin) {
	case "btc":
		walletItem = s.btcService.CreateOneTimeAddress(userId)
	case "ltc":
		walletItem = s.ltcService.CreateOneTimeAddress(userId)
	case "doge":
		walletItem = s.dogeService.CreateOneTimeAddress(userId)
	case "eth":
		walletItem = s.ethService.CreateOneTimeAddress(userId)
	case "trx":
		walletItem = s.trxService.CreateOneTimeAddress(userId)
	case "ton":
		walletItem = s.tonService.CreateOneTimeAddress(userId)
	default:
		err = errors.New("invalid coin name")
	}

	if err != nil {
		log.Panic(err)
		return nil
	}

	return walletItem
}

// DefineBlockchainAndGetCoinBalance -> define a blockchain, init connection
// and get a balance in crypto by address and coin name
func (s *CryptoService) DefineBlockchainAndGetCoinBalance(coin, address string) *pb.CoinBalance {

	var balance *big.Float
	var err error

	switch coin {
	case "btc":
		balance, err = s.btcService.GetBalanceByAddress(address)
	case "ltc":
		balance, err = s.ltcService.GetBalanceByAddress(address)
	case "doge":
		balance, err = s.dogeService.GetBalanceByAddress(address)
	case "ton":
		balance, err = s.tonService.GetBalanceByAddress(address)
	case "eth":
		balance, err = s.ethService.GetBalanceByAddress(address)
	case "trx":
		balance, err = s.trxService.GetBalanceByAddress(address)
	case "sol":
		balance, err = s.solService.GetBalanceByAddress(address) // as a *big.Float
	default:
		err = errors.New("invalid coin name")
	}

	if err != nil {
		log.Panic(err)
		return nil
	}

	return &pb.CoinBalance{CoinName: coin, Balance: balance.String()}
}

func (s *CryptoService) DefineBlockchainAndGetCustomerBalance(currencyType string, customerId int64) *pb.CustomerBalance {

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
func (s *CryptoService) DefineBlockchainAndSendSingleTsx(dto *pb.SendSingleTsxRequest) *pb.TransactionHash {
	var hash string
	var err error

	switch dto.CoinName {
	case "btc":
		hash, err = s.btcService.DefineBlockchainAndSendSingleTransaction(dto)
	case "ltc":
		hash, err = s.ltcService.DefineBlockchainAndSendSingleTransaction(dto)
	case "doge":
		hash, err = s.dogeService.DefineBlockchainAndSendSingleTransaction(dto)
	case "eth":
		hash, err = s.ethService.DefineBlockchainAndSendSingleTransaction(dto)
	case "ton":
		hash, err = s.tonService.DefineBlockchainAndSendSingleTransaction(dto)
	case "trx":
		hash, err = s.trxService.DefineBlockchainAndSendSingleTransaction(dto)
	case "sol":
		hash, err = s.solService.DefineBlockchainAndSendSingleTransaction(dto)
	default:
		err = errors.New("invalid coin name")
	}

	if err != nil {
		log.Panic(err)
		return nil
	}

	return &pb.TransactionHash{TsxHash: hash}
}

// DefineBlockchainAndSendMultiptleTsx -> define a blockchain, init connection
// and send transaction to the list of users by address and coinName
func (s *CryptoService) DefineBlockchainAndSendMultiptleTsx(dto *pb.SendMultipleTsxRequest) *pb.TransactionHash {

	// TODO: impl this method

	var hash string
	var err error

	switch dto.CoinName {
	case "btc":
		hash, err = s.btcService.DefineBlockchainAndSendMultipleTransaction(dto)
	case "ltc":
		hash, err = s.ltcService.DefineBlockchainAndSendMultipleTransaction(dto)
	case "doge":
		hash, err = s.dogeService.DefineBlockchainAndSendMultipleTransaction(dto)
	case "eth":
		hash, err = s.ethService.DefineBlockchainAndSendMultipleTransaction(dto)
	case "ton":
		hash, err = s.tonService.DefineBlockchainAndSendMultipleTransaction(dto)
	case "trx":
		hash, err = s.trxService.DefineBlockchainAndSendMultipleTransaction(dto)
	case "sol":
		hash, err = s.solService.DefineBlockchainAndSendMultipleTransaction(dto)
	default:
		err = errors.New("invalid coin name")
	}

	if err != nil {
		log.Panic(err)
		return nil
	}

	return &pb.TransactionHash{TsxHash: hash}
}
