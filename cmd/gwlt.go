package cmd

import (
	"fmt"
	"wallet-cli/crypto-lib/bitcoin"
	"wallet-cli/crypto-lib/ethereum"
	theopennetwork "wallet-cli/crypto-lib/the-open-network"
	"wallet-cli/crypto-lib/tron"
	"wallet-cli/lib/exceptions"
	"wallet-cli/lib/helpers"

	"github.com/spf13/cobra"
)

// gwltCmd represents the gwlt <get wallet> command
var gwltCmd = &cobra.Command{
	Use:     "get-wallet",
	Aliases: []string{"gwlt"},
	Short:   "Generate a new address.",
	Long: `Generate a new address in the chosen blockchain and
		returns a string address result. There should be a two arguments:
		-> first = a coin name;
		-> second = is a userId.
		
	If something will goes wrong -> an app will throw an error and exit. 
	`,
	Run: func(cmd *cobra.Command, args []string) {

		var coin string
		var userId string
		var address string

		helpers.ValidateArgs(len(args), 2)

		coin = args[0]
		userId = args[1]

		switch coin {
		case "create":
			// run go routine to create a list of all available wallets
			// return a []list of structs -> {coinName: "", address: "", userId: ""}
			// set cache data with generated []list
		case "btc":
			address = bitcoin.CreateWallet(userId)
			// set cache here ? <-
		case "eth":
			address = ethereum.CreateWallet(userId)
		case "trx":
			address = tron.CreateWallet(userId)
		case "ton":
			address = theopennetwork.CreateWallet(userId)
		default:
			exceptions.HandleAnException("Unknown blockchain")
		}

		fmt.Println(address)
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
