package services

import (
	psqlmodels "github.com/mislavperi/fake-instagram-aadbdt/server/internal/infrastructure/psql/models"
	enums "github.com/mislavperi/fake-instagram-aadbdt/server/utils/enums/action"
)

type PlanLogRepository interface {
	InsertPlanChange(userID int64, planID int64) error
	GetUserPlan(userID int64) (*psqlmodels.PlanLog, error)
	InsertAdminPlanChange(userID int64, planID int64) error
}

type PlanLogService struct {
	planLogRepository PlanLogRepository
	logService        LogServiceUpload
}

func NewPlanLogService(planLogRepository PlanLogRepository, logService LogServiceUpload) *PlanLogService {
	return &PlanLogService{
		planLogRepository: planLogRepository,
		logService:        logService,
	}
}

func (s *PlanLogService) InsertPlanChangeLog(userID int64, planID int64) error {
	err := s.planLogRepository.InsertPlanChange(userID, planID)
	if err != nil {
		return err
	}
	s.logService.LogAction(int(userID), enums.GET_CONSUMPTION.String())

	return nil
}

func (s *PlanLogService) InsertAdminPlanChangeLog(userID int64, planID int64) error {
	err := s.planLogRepository.InsertAdminPlanChange(userID, planID)
	if err != nil {
		return err
	}
	s.logService.LogAction(int(userID), enums.GET_CONSUMPTION.String())

	return nil
}

func (s *PlanLogService) GetUserPlan(userID int64) (*psqlmodels.PlanLog, error) {
	planLog, err := s.planLogRepository.GetUserPlan(userID)
	if err != nil {
		return nil, err
	}
	s.logService.LogAction(int(userID), enums.GET_CONSUMPTION.String())

	return planLog, nil
}
