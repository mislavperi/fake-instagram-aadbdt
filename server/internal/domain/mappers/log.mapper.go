package mappers

import (
	psqlmodels "github.com/mislavperi/fake-instagram-aadbdt/server/internal/infrastructure/psql/models"

	"github.com/mislavperi/fake-instagram-aadbdt/server/internal/domain/models"
)

type LogMapper struct {
}

func NewLogMapper() *LogMapper {
	return &LogMapper{}
}

func (m *LogMapper) MapDTOToLogs(logs []psqlmodels.Log) []models.Log {
	var mappedLogs []models.Log
	for _, log := range logs {
		mappedLogs = append(mappedLogs, models.Log{
			ID:        log.ID,
			UserID:    log.UserID,
			Action:    log.Action,
			Timestamp: log.Timestamp,
		})
	}
	return mappedLogs
}
