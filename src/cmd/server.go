package cmd

import (
	"backup-utility/sql"
	"github.com/spf13/cobra"
)

var serverHost string = ""
var serverPort int16 = 0

var serverCMD = &cobra.Command{
	Use:   "server",
	Short: "Manage server list",
	/*PersistentPreRun: func(cmd *cobra.Command, args []string) {

	},*/
}

var serverCMDadd = &cobra.Command{
	Use:   "add",
	Short: "Check connectivity with sync_host",
	/*PersistentPreRun: func(cmd *cobra.Command, args []string) {

	},*/
	Run: func(cmd *cobra.Command, args []string) {
		sql.GetServer()
	},
}
var serverCMDdelete = &cobra.Command{
	Use:   "delete",
	Short: "Check connectivity with sync_host",
	/*PersistentPreRun: func(cmd *cobra.Command, args []string) {

	},*/
}

func init() {

	serverCMD.PersistentFlags().StringVarP(&serverHost, "host", "H", "", "... (required)")
	serverCMD.MarkPersistentFlagRequired("host")
	serverCMD.PersistentFlags().Int16VarP(&serverPort, "port", "P", 0, "... (required)")
	serverCMD.MarkPersistentFlagRequired("port")

	serverCMD.AddCommand(serverCMDadd)
	serverCMD.AddCommand(serverCMDdelete)

	//	rootCmd.AddCommand(serverCMD)
}
