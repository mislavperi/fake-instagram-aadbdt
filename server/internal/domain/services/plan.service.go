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
}

func NewPlanService(planRepository interfaces.PlanRepository, planMapper PlanMapper, logService LogServiceUpload) *PlanService {
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

func (s *PlanService) GetPlan(planName string) (*psqlmodels.Plan, error) {
	plan, err := s.planRepository.GetPlan(planName)
	if err != nil {
		return nil, err
	}
	return plan, nil
}

func (s *PlanService) GetPlanByID(planID int) (*models.Plan, error) {
	plan, err := s.planRepository.GetPlanDetails(planID)
	if err != nil {
		return nil, err
	}
	mappedPlan := s.planMapper.MapPlan(plan)
	return &mappedPlan, nil
}
