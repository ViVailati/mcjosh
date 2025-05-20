package steps

import (
	"fmt"

	"github.com/ViVailati/mcjosh/models"
	"github.com/manifoldco/promptui"
)

const SelectServerTypeStepID = "select_server_type"

type SelectServerTypeStep struct{}

func (s SelectServerTypeStep) ID() string {
	return SelectServerTypeStepID
}

func (s SelectServerTypeStep) Execute(ctx *models.Context, r map[string]string) (*models.StepResult, error) {
	typePrompt := promptui.Select{
		Label: "What's the server type?",
		Items: models.ServerTypes,
		Templates: &promptui.SelectTemplates{
			Active:   `▸ {{ .Name | cyan }}`,
			Inactive: `{{ .Name | faint }}`,
			Selected: `{{ .Name | faint }}`,
		},
	}

	i, _, err := typePrompt.Run()
	if err != nil {
		return nil, fmt.Errorf("%s: %s", SelectServerTypeStepID, err)
	}

	st := models.ServerTypes[i]

	return &models.StepResult{
		Data: map[string]string{
			"server_type_name": st.Name,
			"server_type_url":  st.URL,
		},
		NextStepID:   SelectVanillaVersionStepID,
		ExitWorkflow: false,
	}, nil
}

func (s SelectServerTypeStep) Description() string {
	return "Select the type of Minecraft server"
}
