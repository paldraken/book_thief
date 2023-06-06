package cmd

import (
	"github.com/spf13/cobra"
)

var (
	// Used for flags.
	cfgFile string
	rootCmd = &cobra.Command{
		Use:   "bookthief",
		Short: "Tool to download books from author.today in FB2 format",
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.bookthief.yaml)")
}
