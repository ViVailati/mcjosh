package interfaces

import "github.com/ViVailati/mcjosh/models"

type Step interface {
	ID() string
	Execute(ctx *models.Context, result map[string]string) (*models.StepResult, error)
	Description() string
}
