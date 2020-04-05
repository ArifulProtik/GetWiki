package cli

import (
	"fmt"
	"getwiki/tools"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "help",
	Short: "Use help for know more about gaitonde.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		tools.StartServer()
	},
}

//Execute Executes the Cli
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
