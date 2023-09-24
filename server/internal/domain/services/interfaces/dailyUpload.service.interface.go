package interfaces

import "github.com/mislavperi/fake-instagram-aadbdt/server/internal/domain/models"

type DailyUploadService interface {
	InsertLog(userID int64, pictureID int64, uploadSizeKb uint64) error
	GetStatistics(userID int) (*models.Plan, *uint64, *int, *int, error)
}
