package cmd

import (
	"github.com/jon4hz/ethcli/internal/config"
	"github.com/jon4hz/ethcli/internal/tui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:              "ethcli",
	Short:            "ethcli is a cli tool to work with ethereum",
	TraverseChildren: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return startTui()
	},
}

func init() {
	cobra.OnInitialize(config.Init)
	rootCmd.Flags().StringP("rpc", "r", "http://localhost:8545", "rpc endpoint")
	rootCmd.Flags().StringP("config", "c", "", "config file")

	viper.SetEnvPrefix("ethcli")
	viper.BindPFlag("rpc", rootCmd.Flags().Lookup("rpc"))
	viper.BindPFlag("config", rootCmd.Flags().Lookup("config"))
}

func Execute() error {
	return rootCmd.Execute()
}

func startTui() error {
	return tui.Start()
}
