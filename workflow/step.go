package workflow

type StepResult struct {
	Data         map[string]any
	NextStepID   string
	ExitWorkflow bool
}

type Step interface {
	ID() string
	Execute(ctx *Context) (*StepResult, error)
	Description() string
}
