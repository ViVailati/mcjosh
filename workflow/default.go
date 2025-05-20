package workflow

import (
	client "github.com/ViVailati/mcjosh/http_client"
	"github.com/ViVailati/mcjosh/models"
	"github.com/ViVailati/mcjosh/steps"
)

func DefaultWorkflow() *Workflow {
	w := NewWorkflow(&steps.SelectServerTypeStep{}, &models.Context{
		Client: client.New(),
	})
	w.AddStep(&steps.SelectVanillaVersionStep{})

	return w
}
