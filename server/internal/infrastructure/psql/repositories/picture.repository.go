package repositories

import (
	"mime/multipart"
	"time"

	domainmodels "github.com/mislavperi/fake-instagram-aadbdt/server/internal/domain/models"
	"github.com/mislavperi/fake-instagram-aadbdt/server/internal/infrastructure/psql/models"
	"github.com/mislavperi/fake-instagram-aadbdt/server/utils/errors"
	"gorm.io/gorm"
)

type PictureRepository struct {
	Database *gorm.DB
}

type S3Repository interface {
	UploadToBucket(file multipart.File) (*string, error)
}

func NewPictureRepository(database *gorm.DB) *PictureRepository {
	return &PictureRepository{
		Database: database,
	}
}

func (r *PictureRepository) UploadPicture(title string, description string, hashtags []string, userID int, pictureURI string) (*int64, error) {
	pictureObject := models.Picture{
		UploadDateTime: time.Now(),
		Title:          title,
		Description:    description,
		PictureURI:     pictureURI,
		Hashtags:       hashtags,
		UserID:         int64(userID),
	}
	if err := r.Database.Preload("Plan").Preload("Role").Preload("User").Create(&pictureObject).Error; err != nil {
		return nil, err
	}
	var lastImage models.Picture
	r.Database.Last(&lastImage)
	return &lastImage.ID, nil
}

func (r *PictureRepository) GetImages(filter domainmodels.Filter) ([]*models.Picture, error) {
	var images []*models.Picture
	databaseFilter := r.Database.Preload("User")
	if filter.Title != nil {
		databaseFilter.Where("title = ?", filter.Title)
	}

	if filter.Description != nil {
		databaseFilter.Or("description = ?", filter.Description)
	}

	if filter.TimeRange != nil {
		if filter.TimeRange.Gte != nil && filter.TimeRange.Lte != nil {
			databaseFilter.Where("UploadedDateTime BETWEEN ? AND ?", filter.TimeRange.Gte, filter.TimeRange.Lte)
		} else {
			if filter.TimeRange.Gte != nil {
				databaseFilter.Or("UploadedDateTime >= ?", filter.TimeRange.Gte)
			}
			if filter.TimeRange.Lte != nil {
				databaseFilter.Or("UploadedDateTime =< ?", filter.TimeRange.Lte)
			}
		}
	}
	if filter.Hashtags != nil {
		databaseFilter.Or("hashtags = ?", filter.Hashtags)
	}
	if filter.User != nil {
		databaseFilter.Or("user.username = ?", filter.User)
	}

	result := databaseFilter.Debug().Limit(10).Find(&images)

	if result.Error != nil {
		return nil, result.Error
	}

	return images, nil
}

func (r *PictureRepository) GetUserImages(id int) ([]*models.Picture, error) {
	var images []*models.Picture
	if err := r.Database.Preload("User").Where("user_id = ?", id).Find(&images).Error; err != nil {
		return nil, err
	}

	return images, nil
}

func (r *PictureRepository) GetImageByID(id int) (*models.Picture, error) {
	var image *models.Picture
	if err := r.Database.Where("id = ?", id).First(&image).Error; err != nil {
		return nil, err
	}
	return image, nil
}

func (r *PictureRepository) UpdateImage(id int, description string, hashtags []string, userID int, userRole string) error {
	var image *models.Picture
	if err := r.Database.Preload("User").Where("id = ?", id).First(&image).Error; err != nil {
		return err
	}
	if userRole != "Administrator" {
		if int64(userID) != image.UserID {
			return errors.NewDisallowedResourceError("you don't own the resource you're trying to change")
		}
		image.Description = description
		image.Hashtags = hashtags
		if err := r.Database.Save(&image).Error; err != nil {
			return err
		}
	}
	return nil
}
