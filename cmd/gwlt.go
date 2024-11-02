package cmd

import (
	"fmt"
	cryptolib "wallet-cli/crypto-lib"
	"wallet-cli/lib/helpers"
	"wallet-cli/lib/models"

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

		var walletItem *models.WalletListItem
		var response []*models.WalletListItem

		helpers.ValidateArgs(len(args), 2)
		coin = args[0]
		userId = args[1]

		if coin == "create" {
			response = cryptolib.CreateWalletList(userId)
			fmt.Println("response is: ", response)
		} else {
			walletItem = cryptolib.DefineAndRunBlockchain(&coin, &userId)
		}

		if coin != "create" {
			fmt.Println(walletItem.Address)
		} else {
			fmt.Println("done main_")
		}

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
