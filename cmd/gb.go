package cmd

import (
	"fmt"
	cryptolib "wallet-cli/crypto-lib"
	"wallet-cli/lib/helpers"
	"wallet-cli/lib/models"

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
		var response *models.ResponseBalance

		helpers.ValidateArgs(len(args), 3)
		helpers.CheckAnInternetConnection()

		coin = args[0]
		address = args[1]
		currencyType = args[2]

		response = cryptolib.DefineBlockchainAndGetBalance(coin, address, currencyType)
		fmt.Println(response.CoinName, response.CoinBalance, response.CurrencyType, response.FiatAmount)

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
