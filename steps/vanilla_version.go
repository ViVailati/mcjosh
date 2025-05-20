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

func (s SelectVanillaVersionStep) Execute(ctx *models.Context, result map[string]string) (*models.StepResult, error) {
	stn, sturl, err := getServerTypeInfo(result)
	if err != nil {
		return nil, err
	}

	fmt.Printf("\x1b[33mWe're creating a %s server from: %s\n\x1b[0m", stn, sturl)

	return &models.StepResult{
		Data:         map[string]string{},
		ExitWorkflow: true,
		NextStepID:   GetVanillaVersionsStepID,
	}, nil
}

func getServerTypeInfo(r map[string]string) (string, string, error) {
	stn, ok := r["server_type_name"]
	if !ok {
		return "", "", fmt.Errorf("%s: missing server type name", SelectVanillaVersionStepID)
	}

	sturl, ok := r["server_type_url"]
	if !ok {
		return "", "", fmt.Errorf("%s: missing server type url", SelectVanillaVersionStepID)
	}

	return stn, sturl, nil
}
