package interfaces

import (
	psqlmodels "github.com/mislavperi/fake-instagram-aadbdt/server/internal/infrastructure/psql/models"
)

type PlanRepository interface {
	GetPlans() ([]psqlmodels.Plan, error)
	GetPlan(name string) (*psqlmodels.Plan, error)
}
