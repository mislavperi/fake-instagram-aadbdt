package mappers

import (
	"github.com/mislavperi/fake-instagram-aadbdt/server/internal/domain/models"
	psqlmodels "github.com/mislavperi/fake-instagram-aadbdt/server/internal/infrastructure/psql/models"
)

type UserMapper struct{}

func NewUserMapper() *UserMapper {
	return &UserMapper{}
}

func (m *UserMapper) MapUserToDTO(plan models.Plan) psqlmodels.Plan {
	return psqlmodels.Plan{
		PlanName:          plan.PlanName,
		Cost:              plan.Cost,
		UploadLimitSizeKb: plan.UploadLimitSizeKb,
		DailyUploadLimit:  plan.DailyUploadLimit,
	}
}
