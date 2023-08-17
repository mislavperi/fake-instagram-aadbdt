package services

import (
	"bytes"
	"image"
	"os"
	"strconv"

	"image/jpeg"
	"image/png"
	"mime/multipart"

	"github.com/disintegration/gift"
	"github.com/mislavperi/fake-instagram-aadbdt/server/internal/domain/models"
	"github.com/mislavperi/fake-instagram-aadbdt/server/internal/domain/services/interfaces"
	psqlmodels "github.com/mislavperi/fake-instagram-aadbdt/server/internal/infrastructure/psql/models"
	"github.com/nfnt/resize"
	"golang.org/x/image/bmp"
)

type PictureService struct {
	userService        *UserService
	dailyUploadService interfaces.DailyUploadService

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
	DownloadImage(imageUrl string) ([]byte, error)
}

type PictureRepository interface {
	UploadPicture(string, description string, hashtags []string, userID int, pictureURI string) (*int64, error)
	GetImages(filter models.Filter) ([]*psqlmodels.Picture, error)
	GetUserImages(userID int) ([]*psqlmodels.Picture, error)
	GetImageByID(id int) (*psqlmodels.Picture, error)
	UpdateImage(id int, description string, hashtags []string, userID int, userRole string) error
}

func NewPictureService(userService *UserService, dailyUploadService interfaces.DailyUploadService, pictureRepository PictureRepository, pictureMapper PictureMapper, logRepository interfaces.LogRepository, s3Repository S3Repository) *PictureService {
	return &PictureService{
		pictureRepository: pictureRepository,
		pictureMapper:     pictureMapper,

		dailyUploadService: dailyUploadService,
		userService:        userService,

		logRepository: logRepository,
		s3Repository:  s3Repository,
	}
}

func (s *PictureService) UploadImage(file multipart.File, title string, description string, hashtags []string, id int, height string, width string, fileExt string) error {
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

	newPictureId, err := s.pictureRepository.UploadPicture(title, description, hashtags, id, *pictureURI)
	if err != nil {
		return err
	}

	err = s.dailyUploadService.InsertLog(int64(id), *newPictureId, uint64(len(encodedImage.Bytes())/1024))
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

func (s *PictureService) GetUserImages(id int) ([]models.Picture, error) {
	pictures, err := s.pictureRepository.GetUserImages(id)
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

func (s *PictureService) UpdateImageInformation(imageID int, description string, hashtags []string, userID int) error {
	user, err := s.userService.GetUserInformation(userID)
	if err != nil {
		return err
	}
	err = s.pictureRepository.UpdateImage(imageID, description, hashtags, user.ID, user.Role.Name)
	if err != nil {
		return err
	}
	return nil
}

func (s *PictureService) GetEditedImage(imageID int, height int32, width int32, format string, sepia float32, blur float32) (*bytes.Buffer, error) {
	imageObject, err := s.pictureRepository.GetImageByID(imageID)
	if err != nil {
		return nil, err
	}

	img, err := s.s3Repository.DownloadImage(imageObject.PictureURI)
	if err != nil {
		return nil, err
	}

	decodedImg, _, err := image.Decode(bytes.NewReader(img))
	if err != nil {
		return nil, err
	}

	g := gift.New()

	if height != 0 && width != 0 {
		g.Add(gift.Resize(int(width), int(height), gift.LanczosResampling))
	}

	if sepia != 0 {
		g.Add(gift.Sepia(sepia * 100))
	}

	if blur != 0 {
		g.Add(gift.GaussianBlur(blur))
	}

	dst := image.NewRGBA(g.Bounds(decodedImg.Bounds()))
	var encodedImage bytes.Buffer

	g.Draw(dst, decodedImg)

	switch format {
	case "jpeg":
		err = jpeg.Encode(&encodedImage, dst, &jpeg.Options{})
	case "bmp":
		err = bmp.Encode(&encodedImage, dst)
	case "png":
		err = png.Encode(&encodedImage, dst)
	default:
		err = jpeg.Encode(&encodedImage, dst, &jpeg.Options{})
	}
	if err != nil {
		return nil, err
	}
	os.WriteFile("./image.jpeg", encodedImage.Bytes(), 0777)

	return &encodedImage, nil
}
