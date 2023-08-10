package repositories

import (
	"time"

	"github.com/mislavperi/fake-instagram-aadbdt/server/internal/infrastructure/psql/models"
	"gorm.io/gorm"
)

type LogRepository struct {
	Database *gorm.DB
}

func NewLogRepository(database *gorm.DB) *LogRepository {
	return &LogRepository{
		Database: database,
	}
}

func (r *LogRepository) LogAction(user *models.User, action string) error {
	newLog := models.Log{
		User:      *user,
		Timestamp: time.Now(),
		Action:    action,
	}
	if err := r.Database.Create(&newLog).Error; err != nil {
		return err
	}
	return nil
}
