package services

import (
	"errors"
	"time"

	"github.com/mislavperi/fake-instagram-aadbdt/server/internal/domain/models"
	psqlmodels "github.com/mislavperi/fake-instagram-aadbdt/server/internal/infrastructure/psql/models"
	enums "github.com/mislavperi/fake-instagram-aadbdt/server/utils/enums/action"
)

//go:generate mockery --output=./tests/mocks --name=DailyUploadRepository
type DailyUploadRepository interface {
	InsertLog(userID int64, pictureID int64, uploadSizeKb uint64) error
	GetUserConsumption(userId int) ([]*psqlmodels.DailyUpload, error)
}

//go:generate mockery --output=./tests/mocks --name=UserServiceUpload
type UserServiceUpload interface {
	GetUserInformation(id int) (*models.User, error)
}

//go:generate mockery --output=./tests/mocks --name=PlanLogServiceUpload
type PlanLogServiceUpload interface {
	GetUserPlan(userID int64) (*psqlmodels.PlanLog, error)
}

//go:generate mockery --output=./tests/mocks --name=PlanServiceUpload
type PlanServiceUpload interface {
	GetPlanByID(planID int) (*models.Plan, error)
}

//go:generate mockery --output=./tests/mocks --name=LogServiceUpload
type LogServiceUpload interface {
	LogAction(userID int, action string) error
}

type DailyUploadService struct {
	dailyUploadRepository DailyUploadRepository

	planLogService PlanLogServiceUpload
	planService    PlanServiceUpload
	userService    UserServiceUpload
	logService     LogServiceUpload
}

func NewDailyUploadService(dailyUploadRepository DailyUploadRepository, planLogService PlanLogServiceUpload, planService PlanServiceUpload, userService UserServiceUpload, logService LogServiceUpload) *DailyUploadService {
	return &DailyUploadService{
		dailyUploadRepository: dailyUploadRepository,
		planLogService:        planLogService,
		planService:           planService,
		userService:           userService,
		logService:            logService,
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

	if uint32(len(todayUploads)) >= planLog.Plan.DailyUploadLimit {
		return errors.New("daily upload limit has been reached")
	}
	s.logService.LogAction(userID, enums.GET_CONSUMPTION.String())

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

	mappedPlan, err := s.planService.GetPlanByID(int(planLog.PlanID))
	if err != nil {
		return nil, nil, nil, nil, err
	}

	totalConsumptionCount := len(consumption)
	todayUploadsCount := len(todayUploads)
	s.logService.LogAction(userID, enums.GET_CONSUMPTION.String())

	return mappedPlan, &totalConsumptionKb, &todayUploadsCount, &totalConsumptionCount, nil
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

	mappedPlan, err := s.planService.GetPlanByID(int(planLog.PlanID))
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	totalConsumptionCount := len(consumption)
	todayUploadsCount := len(todayUploads)
	s.logService.LogAction(userID, enums.GET_CONSUMPTION.String())

	return user, mappedPlan, &totalConsumptionKb, &todayUploadsCount, &totalConsumptionCount, nil
}
