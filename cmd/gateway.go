package cmd

import (
	"github.com/spf13/cobra"
	"github.com/zaindeveloper2024/gate-link/internal/gateway"
)

var GatewayCmd = &cobra.Command{
	Use:   "gateway",
	Short: "gateway is a simple API Gateway",
	Long:  "gateway is a simple API Gateway that proxies requests to a target URL",
	Run: func(cmd *cobra.Command, args []string) {
		Init()
		gateway.Handle()
	},
}

func init() {
	RootCmd.AddCommand(GatewayCmd)
}
