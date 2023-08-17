package services

import (
	"github.com/mislavperi/fake-instagram-aadbdt/server/internal/domain/models"
	psqlmodels "github.com/mislavperi/fake-instagram-aadbdt/server/internal/infrastructure/psql/models"
)

type LogRepository interface {
	LogAction(userID int, action string) error
	GetUserLogs(userID int) ([]psqlmodels.Log, error)
}

type LogMapper interface {
	MapDTOToLogs(logs []psqlmodels.Log) []models.Log
}

type LogService struct {
	logRepository LogRepository
	logMapper     LogMapper
}

func NewLogService(logRepository LogRepository, logMapper LogMapper) *LogService {
	return &LogService{
		logRepository: logRepository,
		logMapper:     logMapper,
	}
}

func (s *LogService) LogAction(userID int, action string) error {
	err := s.logRepository.LogAction(userID, action)
	if err != nil {
		return err
	}
	return nil
}

func (s *LogService) GetUserLogs(userID int) ([]models.Log, error) {
	logs, err := s.logRepository.GetUserLogs(userID)
	if err != nil {
		return nil, err
	}
	mappedLogs := s.logMapper.MapDTOToLogs(logs)
	return mappedLogs, nil
}
