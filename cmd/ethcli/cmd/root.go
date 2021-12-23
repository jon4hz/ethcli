package cmd

import (
	"github.com/jon4hz/ethcli/internal/tui"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ethcli",
	Short: "ethcli is a cli tool to work with ethereum",
	RunE: func(cmd *cobra.Command, args []string) error {
		return startTui()
	},
}

func init() {
}

func Execute() error {
	return rootCmd.Execute()
}

func startTui() error {
	return tui.Start()
}
