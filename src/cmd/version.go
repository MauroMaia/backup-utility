package cmd

import (
	"backup-utility/utils"
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Run: func(cmd *cobra.Command, args []string) {
		utils.LogInfo("%s - %s - %s", Version, CommitId, BuildDate)
		fmt.Printf("%s - %s - %s", Version, CommitId, BuildDate)
	},
}
