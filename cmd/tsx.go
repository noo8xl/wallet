package cmd

import (
	"fmt"
	"os"
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
		var tsxDto models.SendTransactionDto

		if len(args) == 0 || len(args) < 4 {
			os.Exit(1)
		} else {
			for _, item := range args {
				fmt.Println("flag is -> ", item)
			}

			tsxDto.CoinName = args[0]
			tsxDto.SenderAddress = args[1]
			tsxDto.RecipientAddress = args[2]
			tsxDto.Amount = args[3]
		}

		fmt.Println("dto ->\n", tsxDto)

		// switch args[0] {
		// case "btc":
		// 	bitcoin.SendSingleBtcTransaction(tsxDto)
		// case "ton":
		// 	theopennetwork.SendSingleTonTransaction(tsxDto)
		// default:
		// 	fmt.Println("Unknown blockchain")
		// }

		fmt.Println("tsx done")
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
