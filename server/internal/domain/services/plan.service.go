package services

import (
	"github.com/mislavperi/fake-instagram-aadbdt/server/internal/domain/models"
	psqlmodels "github.com/mislavperi/fake-instagram-aadbdt/server/internal/infrastructure/psql/models"
)

type PlanRepository interface {
	GetPlans() ([]psqlmodels.Plan, error)
}

type PlanMapper interface {
	MapPlans(plans []psqlmodels.Plan) []models.Plan
}

type PlanService struct {
	planRepository PlanRepository
	planMapper     PlanMapper
}

func NewPlanService(planRepository PlanRepository, planMapper PlanMapper) *PlanService {
	return &PlanService{
		planRepository: planRepository,
		planMapper:     planMapper,
	}
}

func (s *PlanService) GetPlans() ([]models.Plan, error) {
	plans, err := s.planRepository.GetPlans()
	if err != nil {
		return nil, err
	}
	mappedPlans := s.planMapper.MapPlans(plans)
	return mappedPlans, nil
}
