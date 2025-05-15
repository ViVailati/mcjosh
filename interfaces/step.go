package interfaces

import "github.com/ViVailati/mcjosh/models"

type Step interface {
	ID() string
	Execute(ctx *models.Context) (*models.StepResult, error)
	Description() string
}
