package cmd

import (
	"backup-utility/utils"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var Verbose bool
var UseColor bool
var Version string
var CommitId string
var BuildDate string

var rootCmd = &cobra.Command{
	Use:   "backup-utility",
	Short: "backup-utility is a very fast linux backup utility",
	Long: `backup-utility is a very fast linux backup utility.
You should be able to .....

MIT License
Copyright (c) 2021 Mauro Maia
`,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
		os.Exit(1)
	},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		utils.SetLoggerOptions(UseColor, Verbose)
	},
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
	rootCmd.PersistentFlags().BoolVarP(&UseColor, "color", "c", false, "color output")
}

// Execute executes the root command.
func Execute(version string, commitId string, buildDate string) {
	Version = version
	CommitId = commitId
	BuildDate = buildDate

	rootCmd.SetVersionTemplate(fmt.Sprintf("%s - %s - %s", Version, CommitId, BuildDate))

	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		utils.LogError("%s", err.Error())
		os.Exit(1)
	}
}
