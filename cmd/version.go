package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/zaindeveloper2024/gate-link/internal/conf"
)

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Door",
	Long:  "Print the version number of Door and exit",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Door v%s\n", conf.VERSION)
	},
}

func init() {
	RootCmd.AddCommand(VersionCmd)
}
