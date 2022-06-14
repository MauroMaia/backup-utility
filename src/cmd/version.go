package cmd

import (
	"backup-utility/sql"
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
		var migration, _ = sql.GetLasMigration()
		utils.LogInfo("App version: %s - %s - %s", Version, CommitId, BuildDate)
		utils.LogInfo("Database migration: %d", migration)
		fmt.Printf("%s - %s - %s", Version, CommitId, BuildDate)
	},
}
