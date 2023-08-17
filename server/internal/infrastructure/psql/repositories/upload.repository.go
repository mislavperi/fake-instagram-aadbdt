package repositories

import (
	"github.com/mislavperi/fake-instagram-aadbdt/server/internal/infrastructure/psql/models"
	"gorm.io/gorm"
)

type DailyUploadRepository struct {
	Database *gorm.DB
}

func NewDailyUploadRepository(database *gorm.DB) *DailyUploadRepository {
	return &DailyUploadRepository{
		Database: database,
	}
}

func (r *DailyUploadRepository) InsertLog(userID int64, pictureID int64, uploadSizeKb uint64) error {
	if err := r.Database.Preload("User").Preload("Picture").Create(&models.DailyUpload{
		UploadSizeKb: uploadSizeKb,
		PictureID:    pictureID,
		UserID:       userID,
	}).Error; err != nil {
		return err
	}
	return nil
}

func (r *DailyUploadRepository) GetUserConsumption(userId int) ([]*models.DailyUpload, error) {
	var uploads []*models.DailyUpload
	if err := r.Database.Preload("User").Where("user_id = ?", userId).Find(&uploads).Error; err != nil {
		return nil, err
	}
	return uploads, nil
}
