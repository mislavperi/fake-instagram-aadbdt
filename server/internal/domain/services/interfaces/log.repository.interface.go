package interfaces

import psqlmodels "github.com/mislavperi/fake-instagram-aadbdt/server/internal/infrastructure/psql/models"

type LogRepository interface {
	LogAction(userID int, action string) error
	GetUserLogs(userID int) ([]psqlmodels.Log, error)
}
