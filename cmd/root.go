package cmd

import (
	"fmt"
	"os"

	"github.com/ViVailati/mcjosh/workflow"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mcjosh",
	Short: "Minecraft Java (o) Server Helper",
	Long:  `McJosh is a command-line tool for managing Minecraft Java servers.`,
	Run: func(cmd *cobra.Command, args []string) {
		w := workflow.DefaultWorkflow()
		if err := w.Run(); err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(1)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}
