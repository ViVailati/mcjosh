package workflow

import (
	client "github.com/ViVailati/mcjosh/http_client"
	"github.com/ViVailati/mcjosh/models"
	"github.com/ViVailati/mcjosh/steps"
)

func DefaultWorkflow() *Workflow {
	hc := client.New()
	w := NewWorkflow(&steps.SelectServerTypeStep{}, &models.Context{
		Data: map[string]any{
			"client": hc,
		},
	})
	w.AddStep(&steps.SelectVanillaVersionStep{})

	return w
}
