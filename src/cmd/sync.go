package cmd

import (
	"github.com/spf13/cobra"
)

var syncHost string = ""
var syncPort int16 = 0

var syncCMD = &cobra.Command{
	Use:   "sync",
	Short: "Execute a backup operation",
	/*PersistentPreRun: func(cmd *cobra.Command, args []string) {

	},*/
}

var checkConnectivityCMD = &cobra.Command{
	Use:   "check",
	Short: "Check connectivity with sync_host",
	/*PersistentPreRun: func(cmd *cobra.Command, args []string) {

	},*/
}

func init() {

	syncCMD.PersistentFlags().StringVarP(&syncHost, "serverHost", "H", "", "... (required)")
	syncCMD.MarkPersistentFlagRequired("serverHost")
	syncCMD.PersistentFlags().Int16VarP(&syncPort, "serverPort", "P", 0, "... (required)")
	syncCMD.MarkPersistentFlagRequired("serverPort")

	checkConnectivityCMD.PersistentFlags().StringVarP(&syncHost, "serverHost", "H", "", "... (required)")
	checkConnectivityCMD.MarkPersistentFlagRequired("serverHost")
	checkConnectivityCMD.PersistentFlags().Int16VarP(&syncPort, "serverPort", "P", 0, "... (required)")
	checkConnectivityCMD.MarkPersistentFlagRequired("serverPort")

	rootCmd.AddCommand(syncCMD)
	rootCmd.AddCommand(checkConnectivityCMD)
}
