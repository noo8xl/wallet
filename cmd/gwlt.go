package cmd

import (
	"fmt"
	"reflect"
	"strings"
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
		var walletList []*models.WalletListItem

		helpers.ValidateArgs(len(args), 2)
		helpers.CheckAnInternetConnection()

		coin = args[0]
		userId = args[1]

		if coin == "create" {
			walletList = cryptolib.CreateWalletList(&userId)
			var str string

			for _, item := range walletList {
				vals := reflect.ValueOf(*item)
				for i := 0; i < vals.NumField()/2; i++ {
					tmp := strings.Join([]string{vals.Field(i).String(), vals.Field(i + 1).String()}, " ")
					str = strings.Join([]string{str, tmp}, "/")
				}
			}

			fmt.Println(str)
			return
		}

		walletItem = cryptolib.DefineAndRunBlockchain(&coin, &userId)
		fmt.Println(walletItem.Address)

	},
}

func init() {
	rootCmd.AddCommand(gwltCmd)
}
