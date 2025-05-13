package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version information",
	Long:  `Print version information for McJosh.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("McJosh v0.0.1")
	},
}
