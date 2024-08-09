package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/zaindeveloper2024/gate-link/internal/auth"
)

var GenerateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate jwtToken",
	Long:  "Generate jwtToken with the given payload.",
	Run: func(cmd *cobra.Command, args []string) {
		Init()
		token, err := auth.GenerateToken()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error generating token: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Generated JWT token:\n%s\n", token)
	},
}

func init() {
	RootCmd.AddCommand(GenerateCmd)
}
