package cmd

import (
	"fmt"
	cryptolib "wallet-cli/crypto-lib"
	"wallet-cli/lib/helpers"
	"wallet-cli/lib/models"

	"github.com/spf13/cobra"
)

// tsxCmd represents the tsx <transaction> command
var tsxCmd = &cobra.Command{
	Use:   "tsx",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
				and usage of using your command. For example:

				Cobra is a CLI library for Go that empowers applications.
				This application is a tool to generate the needed files
				to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		var hash string
		var tsxDto = new(models.SendTransactionDto)

		helpers.ValidateArgs(len(args), 4)
		helpers.CheckAnInternetConnection()

		tsxDto.CoinName = args[0]
		tsxDto.SenderAddress = args[1]
		tsxDto.RecipientAddress = args[2]
		tsxDto.Amount = args[3]

		hash = cryptolib.DefineBlockchainAndSendTsx(tsxDto)
		fmt.Println(hash)

	},
}

func init() {
	rootCmd.AddCommand(tsxCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// tsxCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// tsxCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
