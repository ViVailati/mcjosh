package steps

import (
	"fmt"

	"github.com/ViVailati/mcjosh/models"
)

const SelectVanillaVersionStepID = "select_vanilla_version"

type SelectVanillaVersionStep struct{}

func (s SelectVanillaVersionStep) ID() string {
	return SelectVanillaVersionStepID
}

func (s SelectVanillaVersionStep) Description() string {
	return "Select the Minecraft version"
}

func (s SelectVanillaVersionStep) Execute(ctx *models.Context) (*models.StepResult, error) {
	fmt.Printf("\x1b[33mWe're creating a %s server from: %s\n\x1b[0m", ctx.Data["server_type"].(models.ServerType).Name, ctx.Data["server_type"].(models.ServerType).URL)
	/*client := http_client.New()
	vs, err := client.GetMinecraftVersions()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	versionPrompt := promptui.Select{
		Label: "What's the Minecraft version?",
		Items: vs,
		Templates: &promptui.SelectTemplates{
			Active:   `▸ {{ .ID | cyan }}`,
			Inactive: `{{ .ID | faint }}`,
			Selected: `{{ .ID | faint }}`,
		},
	}

	i, _, err := versionPrompt.Run()
	if err != nil {
		return nil, fmt.Errorf("%s: %s", SelectVanillaVersionStepID, err)
	}

	return &models.StepResult{
		NextStepID:   SelectServerTypeStepID,
		ExitWorkflow: false,
	}, nil*/

	return &models.StepResult{
		ExitWorkflow: true,
	}, nil
}
