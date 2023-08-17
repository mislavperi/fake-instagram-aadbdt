package services

import (
	"errors"
	"time"

	"github.com/mislavperi/fake-instagram-aadbdt/server/internal/domain/models"
	psqlmodels "github.com/mislavperi/fake-instagram-aadbdt/server/internal/infrastructure/psql/models"
)

type DailyUploadRepository interface {
	InsertLog(userID int64, pictureID int64, uploadSizeKb uint64) error
	GetUserConsumption(userId int) ([]*psqlmodels.DailyUpload, error)
}

type DailyUploadService struct {
	dailyUploadRepository DailyUploadRepository

	planLogService *PlanLogService
	planService    *PlanService
	userService    *UserService
}

func NewDailyUploadService(dailyUploadRepository DailyUploadRepository, planLogService *PlanLogService, planService *PlanService, userService *UserService) *DailyUploadService {
	return &DailyUploadService{
		dailyUploadRepository: dailyUploadRepository,
		planLogService:        planLogService,
		planService:           planService,
		userService:           userService,
	}
}

func (s *DailyUploadService) InsertLog(userID int64, pictureID int64, uploadSizeKb uint64) error {
	err := s.dailyUploadRepository.InsertLog(userID, pictureID, uploadSizeKb)
	if err != nil {
		return err
	}
	return nil
}

func (s *DailyUploadService) GetConsumption(userID int) error {
	var totalConsumptionKb uint64
	var todayUploads []*psqlmodels.DailyUpload
	today := time.Now()

	consumption, err := s.dailyUploadRepository.GetUserConsumption(userID)
	if err != nil {
		return err
	}

	planLog, err := s.planLogService.GetUserPlan(int64(userID))
	if err != nil {
		return err
	}

	for _, entry := range consumption {
		totalConsumptionKb = totalConsumptionKb + entry.UploadSizeKb
		if entry.CreatedAt.Year() == today.Year() && today.Month() == entry.CreatedAt.Month() && entry.CreatedAt.Day() == today.Day() {
			todayUploads = append(todayUploads, entry)
		}
	}

	if totalConsumptionKb >= uint64(planLog.Plan.UploadLimitSizeKb) {
		return errors.New("total consumption limit has been breached")
	}

	if uint32(len(todayUploads)) > planLog.Plan.DailyUploadLimit {
		return errors.New("daily upload limit has been reached")
	}

	return nil
}

func (s *DailyUploadService) GetStatistics(userID int) (*models.Plan, *uint64, *int, *int, error) {
	var totalConsumptionKb uint64
	var todayUploads []*psqlmodels.DailyUpload
	today := time.Now()

	consumption, err := s.dailyUploadRepository.GetUserConsumption(userID)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	planLog, err := s.planLogService.GetUserPlan(int64(userID))
	if err != nil {
		return nil, nil, nil, nil, err
	}

	for _, entry := range consumption {
		totalConsumptionKb = totalConsumptionKb + entry.UploadSizeKb
		if entry.CreatedAt.Year() == today.Year() && today.Month() == entry.CreatedAt.Month() && entry.CreatedAt.Day() == today.Day() {
			todayUploads = append(todayUploads, entry)
		}
	}

	mappedPlan := s.planService.planMapper.MapPlan(&planLog.Plan)

	totalConsumptionCount := len(consumption)
	todayUploadsCount := len(todayUploads)

	return &mappedPlan, &totalConsumptionKb, &todayUploadsCount, &totalConsumptionCount, nil
}

func (s *DailyUploadService) GetExpandedStatistics(userID int) (*models.User, *models.Plan, *uint64, *int, *int, error) {
	var totalConsumptionKb uint64
	var todayUploads []*psqlmodels.DailyUpload
	today := time.Now()

	consumption, err := s.dailyUploadRepository.GetUserConsumption(userID)
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	planLog, err := s.planLogService.GetUserPlan(int64(userID))
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	user, err := s.userService.GetUserInformation(userID)
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	for _, entry := range consumption {
		totalConsumptionKb = totalConsumptionKb + entry.UploadSizeKb
		if entry.CreatedAt.Year() == today.Year() && today.Month() == entry.CreatedAt.Month() && entry.CreatedAt.Day() == today.Day() {
			todayUploads = append(todayUploads, entry)
		}
	}

	mappedPlan := s.planService.planMapper.MapPlan(&planLog.Plan)

	totalConsumptionCount := len(consumption)
	todayUploadsCount := len(todayUploads)

	return user, &mappedPlan, &totalConsumptionKb, &todayUploadsCount, &totalConsumptionCount, nil
}
