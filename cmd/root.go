package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/zaindeveloper2024/gate-link/cmd/flags"
)

var RootCmd = &cobra.Command{
	Use:   "gate-link",
	Short: "gate-link is a simple API Gateway",
	Long:  "gate-link is a simple API Gateway that proxies requests to a target URL",
}

func Main() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	RootCmd.PersistentFlags().BoolVar(&flags.Debug, "debug", false, "Enable debug mode")
	RootCmd.PersistentFlags().IntVar(&flags.Port, "port", 8080, "Port to run the API Gateway on")
	RootCmd.PersistentFlags().BoolVar(&flags.Auth, "auth", false, "Enable authentication")
}
