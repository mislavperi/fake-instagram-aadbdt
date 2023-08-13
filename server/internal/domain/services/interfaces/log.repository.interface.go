package interfaces

import (
	psqlmodels "github.com/mislavperi/fake-instagram-aadbdt/server/internal/infrastructure/psql/models"
)

type LogRepository interface {
	LogAction(user *psqlmodels.User, action string) error
}
