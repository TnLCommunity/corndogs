package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = NewRootCommand()

func NewRootCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "",
		Short: "The corndogs cli",
		Long:  "The corndogs cli. See available commands below.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return fmt.Errorf("This is the root command. It doesnt do anything without a subcommand listed below.")
		},
	}
	return rootCmd
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
