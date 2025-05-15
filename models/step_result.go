package models

type StepResult struct {
	Data         map[string]any
	NextStepID   string
	ExitWorkflow bool
}
