package workflow

import (
	"fmt"

	"github.com/ViVailati/mcjosh/http_client"
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
		ctx = &models.Context{
			Client: http_client.New(),
		}
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
	stepResults := make(map[string]string)

	for {
		step, ok := w.Steps[w.CurrentStep]
		if !ok {
			return fmt.Errorf("step %s not in workflow", w.CurrentStep)
		}

		result, err := step.Execute(w.Context, stepResults)
		if err != nil {
			return err
		}

		if result.ExitWorkflow {
			break
		}

		stepResults = result.Data
		w.CurrentStep = result.NextStepID
	}

	return nil
}
