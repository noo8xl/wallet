package cmd

import (
	"fmt"
	"os"
	"wallet-cli/crypto-lib/bitcoin"
	theopennetwork "wallet-cli/crypto-lib/the-open-network"

	"github.com/spf13/cobra"
)

// gwltCmd represents the gwlt command
var gwltCmd = &cobra.Command{
	Use:   "gwlt",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		var flag string
		var testId string = "userId"

		if len(args) == 0 {
			os.Exit(1)
		} else {
			for _, item := range args {
				fmt.Println("flag is -> ", item)
			}

			flag = args[0]
		}

		switch flag {
		case "btc":
			bitcoin.CreateWallet(testId)
		case "ton":
			theopennetwork.CreateWallet(testId)
		default:
			fmt.Println("Unknown blockchain")
		}

		fmt.Println("gwlt done")
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
