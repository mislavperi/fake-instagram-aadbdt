package services

import (
	"mime/multipart"

	"github.com/mislavperi/fake-instagram-aadbdt/server/internal/domain/models"
	"github.com/mislavperi/fake-instagram-aadbdt/server/internal/domain/services/interfaces"
	psqlmodels "github.com/mislavperi/fake-instagram-aadbdt/server/internal/infrastructure/psql/models"
)

type PictureService struct {
	pictureRepository PictureRepository
	logRepository     interfaces.LogRepository
}

type PictureRepository interface {
	UploadPicture(file multipart.File, string, description string, hashtags []string, user psqlmodels.User) error
}

func NewPictureService(pictureRepository PictureRepository, logRepository interfaces.LogRepository) *PictureService {
	return &PictureService{
		pictureRepository: pictureRepository,
		logRepository:     logRepository,
	}
}

func (s *PictureService) UploadImage(file multipart.File, title string, description string, hashtags []string, user models.User) error {
	mappedUser := psqlmodels.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
		Email:     user.Email,
		Role: psqlmodels.Role{
			Name: user.Role.Name,
		},
		Plan: psqlmodels.Plan{
			PlanName:          user.Plan.PlanName,
			UploadLimitSizeKb: user.Plan.UploadLimitSizeKb,
			DailyUploadLimit:  user.Plan.DailyUploadLimit,
			Cost:              user.Plan.Cost,
		},
	}

	err := s.pictureRepository.UploadPicture(file, title, description, hashtags, mappedUser)
	if err != nil {
		return err
	}
	return nil
}
