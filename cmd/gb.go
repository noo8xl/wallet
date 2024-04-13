package cmd

import (
	"fmt"
	"log"
	"math/big"
	"os"
	"wallet-cli/api"
	"wallet-cli/crypto-lib/bitcoin"
	theopennetwork "wallet-cli/crypto-lib/the-open-network"

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

		var flag string
		var address string
		var balance *big.Float
		var fiatBalance float64
		f := new(big.Float)

		if len(args) == 0 {
			os.Exit(1)
		} else {

			flag = args[0]
			address = args[1]

			fmt.Println("coin name is -> ", flag)
			fmt.Println("coin address is -> ", address)

		}

		switch flag {
		case "btc":
			balance = bitcoin.GetBitcoinAddressBalance(address)
			fiatBalance = api.GetRate("bitcoin", "usd")
		case "ton":
			balance = theopennetwork.GetTonBalanceByAddress(address)
			fiatBalance = api.GetRate("the-open-network", "usd")
		default:
			fmt.Println("Unknown blockchain")
		}

		f.SetFloat64(fiatBalance)
		balInUsd := new(big.Float).SetPrec(20).Mul(f, balance)

		log.Println(flag, "balance is ->", balance, "= $", balInUsd)
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
