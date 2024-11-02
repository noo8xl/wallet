package cryptolib

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"wallet-cli/config"
	"wallet-cli/crypto-lib/bitcoin"
	theopennetwork "wallet-cli/crypto-lib/the-open-network"
	"wallet-cli/crypto-lib/tron"
	"wallet-cli/lib/exceptions"
	"wallet-cli/lib/models"
)

type routineOpts struct {
	userId   string
	coinName string
	result   chan<- *models.WalletListItem
	wg       *sync.WaitGroup
}

func CreateWalletList(userId string) []*models.WalletListItem {

	result := make(chan *models.WalletListItem, 4)
	var walletItem *models.WalletListItem
	var wg sync.WaitGroup
	coinList := config.GetAvailableCoinList()
	var walletList []*models.WalletListItem

	for _, item := range coinList {
		wg.Add(1)
		opt := &routineOpts{
			userId:   userId,
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

	fmt.Println("done ------")

	for _, item := range walletList {
		fmt.Println(" wallet item is -> ", item)
	}

	return walletList
}

func worker(opts *routineOpts) {
	defer opts.wg.Done()
	fmt.Printf("worker coinName: %s \n", opts.coinName)
	opts.result <- DefineAndRunBlockchain(&opts.coinName, &opts.userId)
	// opts.result <- &models.WalletListItem{}

}

func DefineAndRunBlockchain(coin, userId *string) *models.WalletListItem {
	switch strings.ToLower(*coin) {
	case "btc":
		return bitcoin.CreateWallet(userId)
	case "eth":
		// return ethereum.CreateWallet(userId)
		return &models.WalletListItem{CoinName: "eth", Address: "will be"}
	case "trx":
		return tron.CreateWallet(userId)
	case "ton":
		return theopennetwork.CreateWallet(userId)
	default:
		exceptions.HandleAnException("Unknown blockchain")
		os.Exit(1)
	}
	return nil
}
