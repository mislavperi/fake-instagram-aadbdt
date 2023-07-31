package mappers

import (
	"github.com/mislavperi/fake-instagram-aadbdt/server/internal/domain/models"
	psqlmodels "github.com/mislavperi/fake-instagram-aadbdt/server/internal/infrastructure/psql/models"
)

type PlanMapper struct{}

func NewPlanMapper() *PlanMapper {
	return &PlanMapper{}
}

func (m *PlanMapper) MapPlans(plans []psqlmodels.Plan) []models.Plan {
	var mappedPlans []models.Plan

	for _, plan := range plans {
		mappedPlans = append(mappedPlans, models.Plan{
			PlanName:          plan.PlanName,
			DailyUploadLimit:  plan.DailyUploadLimit,
			UploadLimitSizeKb: plan.UploadLimitSizeKb,
			Cost:              plan.Cost,
		})
	}
	return mappedPlans
}
