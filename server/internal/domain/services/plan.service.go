package services

import (
	"github.com/mislavperi/fake-instagram-aadbdt/server/internal/domain/models"
	"github.com/mislavperi/fake-instagram-aadbdt/server/internal/domain/services/interfaces"
	psqlmodels "github.com/mislavperi/fake-instagram-aadbdt/server/internal/infrastructure/psql/models"
)

type PlanMapper interface {
	MapPlans(plans []psqlmodels.Plan) []models.Plan
	MapPlan(plan *psqlmodels.Plan) models.Plan
}

type PlanService struct {
	planRepository interfaces.PlanRepository
	planMapper     PlanMapper
	logRepository  interfaces.LogRepository
}

func NewPlanService(planRepository interfaces.PlanRepository, planMapper PlanMapper, logRepository interfaces.LogRepository) *PlanService {
	return &PlanService{
		planRepository: planRepository,
		planMapper:     planMapper,
		logRepository:  logRepository,
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

func (s *PlanService) GetPlan(planName string) (*psqlmodels.Plan, error) {
	plan, err := s.planRepository.GetPlan(planName)
	if err != nil {
		return nil, err
	}
	return plan, nil
}
