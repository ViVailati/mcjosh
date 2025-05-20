package models

type StepResult struct {
	Data         map[string]string
	NextStepID   string
	ExitWorkflow bool
}
