package services

import psqlmodels "github.com/mislavperi/fake-instagram-aadbdt/server/internal/infrastructure/psql/models"

type PlanLogRepository interface {
	InsertPlanChange(userID int64, planID int64) error
	GetUserPlan(userID int64) (*psqlmodels.PlanLog, error)
	InsertAdminPlanChange(userID int64, planID int64) error
}

type PlanLogService struct {
	planLogRepository PlanLogRepository
}

func NewPlanLogService(planLogRepository PlanLogRepository) *PlanLogService {
	return &PlanLogService{
		planLogRepository: planLogRepository,
	}
}

func (s *PlanLogService) InsertPlanChangeLog(userID int64, planID int64) error {
	err := s.planLogRepository.InsertPlanChange(userID, planID)
	if err != nil {
		return err
	}
	return nil
}

func (s *PlanLogService) InsertAdminPlanChangeLog(userID int64, planID int64) error {
	err := s.planLogRepository.InsertAdminPlanChange(userID, planID)
	if err != nil {
		return err
	}
	return nil
}

func (s *PlanLogService) GetUserPlan(userID int64) (*psqlmodels.PlanLog, error) {
	planLog, err := s.planLogRepository.GetUserPlan(userID)
	if err != nil {
		return nil, err
	}
	return planLog, nil
}
