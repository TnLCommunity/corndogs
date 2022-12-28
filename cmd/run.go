package cmd

import (
	"github.com/TnLCommunity/corndogs/server"
	"github.com/spf13/cobra"
)

var runCmd = NewRunCommand()

func NewRunCommand() *cobra.Command {
	runCmd := &cobra.Command{
		Use:   "run",
		Short: "Run the corndogs service",
		Long:  "Run the corndogs service",
		Run: func(cmd *cobra.Command, args []string) {
			server.SetupAndRun()
		},
	}

	rootCmd.AddCommand(runCmd)
	return runCmd
}
