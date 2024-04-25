package cmd

import (
	"fmt"
	"log"
	"os"
	"time"
	"wallet-cli/crypto-lib/bitcoin"
	"wallet-cli/crypto-lib/ethereum"
	theopennetwork "wallet-cli/crypto-lib/the-open-network"
	"wallet-cli/lib/cache"

	"github.com/spf13/cobra"
)

// gwltCmd represents the gwlt <get wallet> command
var gwltCmd = &cobra.Command{
	Use:   "gwlt",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
				and usage of using your command. For example:

				Cobra is a CLI library for Go that empowers applications.
				This application is a tool to generate the needed files
				to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		stamp := time.Now().UnixMilli()

		var flag string
		var userId string
		var address string

		if len(args) == 0 {
			os.Exit(1)
		} else {

			flag = args[0]
			userId = args[1]

			fmt.Println("coin name is -> ", flag)
			fmt.Println("userId is -> ", userId)
		}

		switch flag {
		case "create":
			// run go routine to create a list of all available wallets
			// return a []list of structs -> {coinName: "", address: "", userId: ""}
			// set cache data with generated []list
		case "btc":
			address = bitcoin.CreateWallet(userId)
			// set cache val to have an access from the main API
			cache.SetWalletData(userId, flag, address)
		case "eth":
			address = ethereum.CreateWallet(userId)
			// set cache val to have an access from the main API
			cache.SetWalletData(userId, flag, address)
		case "ton":
			address = theopennetwork.CreateWallet(userId)
			// set cache val to have an access from the main API
			cache.SetWalletData(userId, flag, address)
		default:
			fmt.Println("Unknown blockchain")
		}

		execTime := time.Now().UnixMilli() - stamp
		log.Println("ETH address is -> ", address)
		fmt.Println("gwlt done in ", execTime, "ms")
	},
}

func init() {
	rootCmd.AddCommand(gwltCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// gwltCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// gwltCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
