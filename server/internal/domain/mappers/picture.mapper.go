package mappers

import (
	"github.com/mislavperi/fake-instagram-aadbdt/server/internal/domain/models"
	psqlmodels "github.com/mislavperi/fake-instagram-aadbdt/server/internal/infrastructure/psql/models"
)

type PictureMapper struct {
}

func NewPictureMapper() *PictureMapper {
	return &PictureMapper{}
}

func (m *PictureMapper) MapDTOToPictures(pictures []*psqlmodels.Picture) []models.Picture {
	var mappedPictures []models.Picture
	for _, picture := range pictures {
		mappedPictures = append(mappedPictures, models.Picture{
			ID:             picture.ID,
			Title:          picture.Title,
			Description:    picture.Description,
			PictureURI:     picture.PictureURI,
			UploadDateTime: picture.UploadDateTime,
			Hashtags:       picture.Hashtags,
			User: models.User{
				Username: picture.User.Username,
			},
		})
	}

	return mappedPictures
}

func (m *PictureMapper) MapDTOToPicture(picture *psqlmodels.Picture) models.Picture {
	mappedPicture := models.Picture{
		ID:          picture.ID,
		Title:       picture.Title,
		Description: picture.Description,
		PictureURI:  picture.PictureURI,
		Hashtags:    picture.Hashtags,
	}

	return mappedPicture
}
