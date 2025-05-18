package models

type StepResult struct {
	Data         any
	NextStepID   string
	ExitWorkflow bool
}
