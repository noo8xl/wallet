package cryptolib

import (
	"math/big"
	"sync"
	"wallet-cli/lib/models"
)

type routineOpts struct {
	userId   string
	coinName string
	result   chan<- *models.WalletListItem
	wg       *sync.WaitGroup
}

// getaBigFloat -> convert a float64 fiat value to the *big.Float
// to get a current amount of the coins
func getaBigFloat(fiat float64, crypto *big.Float) *big.Float {
	f := new(big.Float)
	f.SetFloat64(fiat)
	return new(big.Float).SetPrec(20).Mul(f, crypto)
}

// worker -> will be spanned in the create function to
// create a list of wallets for the user
func worker(opts *routineOpts) {
	defer opts.wg.Done()
	opts.result <- DefineAndRunBlockchain(&opts.coinName, &opts.userId)
}
