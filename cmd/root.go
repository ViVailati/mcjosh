package cmd

import (
	"fmt"
	"os"

	"github.com/ViVailati/mcjosh/models"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mcjosh",
	Short: "Minecraft Java (o) Server Helper",
	Long:  `McJosh is a command-line tool for managing Minecraft Java servers.`,
	Run: func(cmd *cobra.Command, args []string) {
		prompt := promptui.Select{
			Label: "What's the server type?",
			Items: models.ServerTypes,
			Templates: &promptui.SelectTemplates{
				Active:   `▸ {{ .Name | cyan }}`,
				Inactive: `{{ .Name | faint }}`,
			},
		}

		i, _, err := prompt.Run()
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		fmt.Printf("We're creating a %s server from: %s\n", models.ServerTypes[i].Name, models.ServerTypes[i].URL)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}
