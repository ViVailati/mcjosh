package cmd

import (
	"fmt"
	"os"

	client "github.com/ViVailati/mcjosh/http_client"
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

		st := models.ServerTypes[i]

		fmt.Printf("We're creating a %s server from: %s\n", st.Name, st.URL)

		client := client.New()
		if err := client.GetMinecraftVersion(); err != nil {
			fmt.Println("Error:", err)
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
