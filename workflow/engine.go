package workflow

import (
	"fmt"

	"github.com/ViVailati/mcjosh/interfaces"
	"github.com/ViVailati/mcjosh/models"
)

type Workflow struct {
	Steps       map[string]interfaces.Step
	InitialStep string
	CurrentStep string
	Context     *models.Context
}

func NewWorkflow(initialStep interfaces.Step, ctx *models.Context) *Workflow {
	if ctx == nil {
		ctx = &models.Context{Data: make(map[string]any)}
	}

	steps := make(map[string]interfaces.Step)
	steps[initialStep.ID()] = initialStep
	return &Workflow{
		Steps:       steps,
		InitialStep: initialStep.ID(),
		CurrentStep: initialStep.ID(),
		Context:     ctx,
	}
}

func (w *Workflow) AddStep(step interfaces.Step) {
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
