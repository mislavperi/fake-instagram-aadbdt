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

func (r *LogRepository) LogAction(userID int, action string) error {
	newLog := models.Log{
		UserID:    int64(userID),
		Timestamp: time.Now(),
		Action:    action,
	}
	if err := r.Database.Create(&newLog).Error; err != nil {
		return err
	}
	return nil
}

func (r *LogRepository) GetUserLogs(userID int) ([]models.Log, error) {
	var userLogs []models.Log
	if err := r.Database.Where("user_id = ?", userID).Limit(15).Find(&userLogs).Error; err != nil {
		return nil, err
	}
	return userLogs, nil
}
