// Cli Package for Using Cli Utilities For Better Deployment.

package cli

import (
	"fmt"
	"getwiki/tools"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(migrateCmd, webCmd)
}

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrates All Latest Model To Database.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Not Implemented")
	},
}
var webCmd = &cobra.Command{
	Use:   "web",
	Short: "Web Command Start The Web Api.",
	Run: func(cmd *cobra.Command, args []string) {
		tools.StartServer()
	},
}
