package repositories

import (
	"mime/multipart"
	"time"

	"github.com/mislavperi/fake-instagram-aadbdt/server/internal/infrastructure/psql/models"
	"gorm.io/gorm"
)

type PictureRepository struct {
	Database     *gorm.DB
	s3Repository S3Repository
}

type S3Repository interface {
	UploadToBucket(file multipart.File) (*string, error)
}

func NewPictureRepository(database *gorm.DB, s3Repository S3Repository) *PictureRepository {
	return &PictureRepository{
		Database:     database,
		s3Repository: s3Repository,
	}
}

func (r *PictureRepository) UploadPicture(file multipart.File, title string, description string, hashtags []string, user models.User) error {
	pictureURI, err := r.s3Repository.UploadToBucket(file)
	if err != nil {
		return err
	}
	pictureObject := models.Picture{
		UploadDateTime: time.Now(),
		Title:          title,
		Description:    description,
		PictureURI:     *pictureURI,
		Hashtags:       hashtags,
		User:           user,
	}
	if err := r.Database.Create(&pictureObject).Error; err != nil {
		return err
	}
	return nil
}
