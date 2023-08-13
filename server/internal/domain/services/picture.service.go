package services

import (
	"bytes"
	"image"
	"strconv"

	"image/jpeg"
	"image/png"
	"mime/multipart"

	"github.com/mislavperi/fake-instagram-aadbdt/server/internal/domain/models"
	"github.com/mislavperi/fake-instagram-aadbdt/server/internal/domain/services/interfaces"
	psqlmodels "github.com/mislavperi/fake-instagram-aadbdt/server/internal/infrastructure/psql/models"
	"github.com/nfnt/resize"
	"golang.org/x/image/bmp"
)

type PictureService struct {
	userService *UserService

	pictureRepository PictureRepository
	pictureMapper     PictureMapper

	s3Repository  S3Repository
	logRepository interfaces.LogRepository
}

type PictureMapper interface {
	MapDTOToPictures(pictures []*psqlmodels.Picture) []models.Picture
	MapDTOToPicture(picture *psqlmodels.Picture) models.Picture
}
type S3Repository interface {
	UploadToBucket(file bytes.Buffer, fileExt string) (*string, error)
}

type PictureRepository interface {
	UploadPicture(string, description string, hashtags []string, user psqlmodels.User, pictureURI string) error
	GetImages(filter models.Filter) ([]*psqlmodels.Picture, error)
	GetUserImages(userID int) ([]*psqlmodels.Picture, error)
	GetImageByID(id int) (*psqlmodels.Picture, error)
	UpdateImage(id int, description string, hashtags []string, userID int, userRole string) error
}

func NewPictureService(userService *UserService, pictureRepository PictureRepository, pictureMapper PictureMapper, logRepository interfaces.LogRepository, s3Repository S3Repository) *PictureService {
	return &PictureService{
		pictureRepository: pictureRepository,
		pictureMapper:     pictureMapper,

		userService: userService,

		logRepository: logRepository,
		s3Repository:  s3Repository,
	}
}

func (s *PictureService) UploadImage(file multipart.File, title string, description string, hashtags []string, username string, height string, width string, fileExt string) error {
	user, err := s.userService.GetUserInformation(username)
	if err != nil {
		return err
	}
	mappedUser := s.userService.MapUserToDTO(*user)
	decodedImage, _, err := image.Decode(file)
	if err != nil {
		return err
	}
	widthInt, err := strconv.Atoi(width)
	if err != nil {
		return err
	}

	heightInt, err := strconv.Atoi(height)
	if err != nil {
		return err
	}
	var resizedImage image.Image
	resizedImage = decodedImage
	if widthInt != 0 && heightInt != 0 {
		resizedImage = resize.Resize(uint(widthInt), uint(heightInt), decodedImage, resize.Lanczos3)
	}
	var encodedImage bytes.Buffer

	switch fileExt {
	case "jpeg":
		err = jpeg.Encode(&encodedImage, resizedImage, &jpeg.Options{})
	case "bmp":
		err = bmp.Encode(&encodedImage, resizedImage)
	case "png":
		err = png.Encode(&encodedImage, resizedImage)
	default:
		err = jpeg.Encode(&encodedImage, resizedImage, &jpeg.Options{})
	}
	if err != nil {
		return err
	}

	pictureURI, err := s.s3Repository.UploadToBucket(encodedImage, fileExt)
	if err != nil {
		return err
	}

	err = s.pictureRepository.UploadPicture(title, description, hashtags, mappedUser, *pictureURI)
	if err != nil {
		return err
	}
	return nil
}

func (s *PictureService) GetImages(filter models.Filter) ([]models.Picture, error) {
	pictures, err := s.pictureRepository.GetImages(filter)
	if err != nil {
		return nil, err
	}

	mappedPictures := s.pictureMapper.MapDTOToPictures(pictures)
	return mappedPictures, nil
}

func (s *PictureService) GetUserImages(username string) ([]models.Picture, error) {
	user, err := s.userService.GetUserInformation(username)
	if err != nil {
		return nil, err
	}

	pictures, err := s.pictureRepository.GetUserImages(user.ID)
	if err != nil {
		return nil, err
	}
	mappedPictures := s.pictureMapper.MapDTOToPictures(pictures)
	return mappedPictures, nil
}

func (s *PictureService) GetImageByID(id int) (*models.Picture, error) {
	image, err := s.pictureRepository.GetImageByID(id)
	if err != nil {
		return nil, err
	}
	mappedImage := s.pictureMapper.MapDTOToPicture(image)
	return &mappedImage, nil
}

func (s *PictureService) UpdateImageInformation(imageID int, description string, hashtags []string, username string) error {
	user, err := s.userService.GetUserInformation(username)
	if err != nil {
		return err
	}
	err = s.pictureRepository.UpdateImage(imageID, description, hashtags, user.ID, user.Role.Name)
	if err != nil {
		return err
	}
	return nil
}
