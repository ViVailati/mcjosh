package workflow

import "fmt"

type Workflow struct {
	Steps       map[string]Step
	InitialStep string
	CurrentStep string
	Context     *Context
}

func NewWorkflow(initialStep string) *Workflow {
	return &Workflow{
		Steps:       make(map[string]Step),
		InitialStep: initialStep,
		CurrentStep: initialStep,
		Context:     &Context{Data: make(map[string]any)},
	}
}

func (w *Workflow) AddStep(step Step) {
	w.Steps[step.ID()] = step
}

func (w *Workflow) Run() error {
	w.CurrentStep = w.InitialStep

	for {
		step, ok := w.Steps[w.CurrentStep]
		if !ok {
			return fmt.Errorf("step %s not in workflow", w.CurrentStep)
		}

		result, err := step.Execute(w.Context)
		if err != nil {
			return err
		}

		if result.ExitWorkflow {
			break
		}

		w.CurrentStep = result.NextStepID
	}

	return nil
}
