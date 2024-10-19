package cmd

import (
	"fmt"
	"log"
	"math/big"
	"os"
	"time"
	"wallet-cli/api"
	"wallet-cli/crypto-lib/bitcoin"
	"wallet-cli/crypto-lib/ethereum"
	theopennetwork "wallet-cli/crypto-lib/the-open-network"
	"wallet-cli/lib/exceptions"

	"github.com/spf13/cobra"
)

// gbCmd represents the gb <get balance> command
var gbCmd = &cobra.Command{
	Use:   "gb",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
				and usage of using your command. For example:

				Cobra is a CLI library for Go that empowers applications.
				This application is a tool to generate the needed files
				to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		var coin string
		var address string
		var currencyType string
		var balance *big.Float
		var fiatBalance float64
		f := new(big.Float)
		stamp := time.Now().UnixMilli()

		if len(args) == 0 {
			os.Exit(1)
		} else {

			coin = args[0]
			address = args[1]
			currencyType = args[2]

			// fmt.Println("coin name is -> ", coin)
			// fmt.Println("coin address is -> ", address)
			// fmt.Println("currency type is -> ", currencyType)

		}

		switch coin {
		case "btc":
			balance = bitcoin.GetBitcoinAddressBalance(address)
			// balance =
			fiatBalance = api.GetRate("bitcoin", currencyType)
		case "ton":
			balance = theopennetwork.GetTonBalanceByAddress(address)
			fiatBalance = api.GetRate("the-open-network", currencyType)
		case "eth":
			balance = ethereum.GetEthBalanceByAddress(address)
			fiatBalance = api.GetRate("ethereum", currencyType)
		case "trx":
			balance = theopennetwork.GetTonBalanceByAddress(address)
			fiatBalance = api.GetRate("tron", currencyType)
		default:
			exceptions.HandleAnException("Unknown blockchain")
		}

		f.SetFloat64(fiatBalance)
		balInUsd := new(big.Float).SetPrec(20).Mul(f, balance)

		execTime := time.Now().UnixMilli() - stamp
		log.Println(coin, "balance is ->", balance, "=", currencyType, ":", balInUsd)
		fmt.Println("gb done in ", execTime, "ms")
	},
}

func init() {
	rootCmd.AddCommand(gbCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// gbCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// gbCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
