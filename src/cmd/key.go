package cmd

import (
	"github.com/spf13/cobra"
)

var keyCMD = &cobra.Command{
	Use:   "key",
	Short: "List encryption keys when they exist",
	/*PersistentPreRun: func(cmd *cobra.Command, args []string) {

	},*/
}

var importKeyCMD = &cobra.Command{
	Use:   "import",
	Short: "import encryption keys when they exist",
	/*PersistentPreRun: func(cmd *cobra.Command, args []string) {

	},*/
}

var exportKeyCMD = &cobra.Command{
	Use:   "export",
	Short: "export encryption keys when they exist",
	/*PersistentPreRun: func(cmd *cobra.Command, args []string) {

	},*/
}

func init() {
	rootCmd.AddCommand(keyCMD)
	rootCmd.AddCommand(exportKeyCMD)
	rootCmd.AddCommand(importKeyCMD)
}
