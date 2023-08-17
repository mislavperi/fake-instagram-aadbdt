package controllers

import (
	"bytes"
	"encoding/json"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mislavperi/fake-instagram-aadbdt/server/internal/domain/models"
)

type PictureController struct {
	pictureService PictureService
}

type PictureService interface {
	UploadImage(file multipart.File, title string, description string, hashtags []string, userID int, height string, width string, fileExt string) error
	GetImages(filter models.Filter) ([]models.Picture, error)
	GetUserImages(userID int) ([]models.Picture, error)
	GetImageByID(id int) (*models.Picture, error)
	UpdateImageInformation(imageID int, description string, hashtags []string, userID int) error
	GetEditedImage(imageID int, height int32, width int32, format string, sepia float32, blur float32) (*bytes.Buffer, error)
}

func NewPictureController(pictureService PictureService) *PictureController {
	return &PictureController{
		pictureService: pictureService,
	}
}

func (c *PictureController) UploadImage() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var pictureUpload models.PictureUpload
		ctx.ShouldBind(&pictureUpload)
		var parsedHashtags []string
		if err := json.Unmarshal([]byte(pictureUpload.Hashtags), &parsedHashtags); err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
		}
		file, err := pictureUpload.Picture.Open()
		if err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
		}
		identifier, err := strconv.Atoi(ctx.GetHeader("Identifier"))
		if err != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}

		c.pictureService.UploadImage(file, pictureUpload.Title, pictureUpload.Description, parsedHashtags, identifier, pictureUpload.Height, pictureUpload.Width, pictureUpload.Format)
		ctx.JSON(http.StatusOK, nil)
	}
}

func (c *PictureController) GetImages() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		clientFilter := models.Filter{}
		requestFilter := ctx.Query("filter")
		err := json.Unmarshal([]byte(requestFilter), &clientFilter)
		if err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		mappedPictures, err := c.pictureService.GetImages(clientFilter)
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		ctx.JSON(http.StatusOK, mappedPictures)
	}
}

func (c *PictureController) GetUserImages() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		identifier, err := strconv.Atoi(ctx.GetHeader("Identifier"))
		if err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		mappedPictures, err := c.pictureService.GetUserImages(identifier)
		if err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
		}
		ctx.JSON(http.StatusOK, mappedPictures)
	}
}

func (c *PictureController) GetSpecificUserImages() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var userID int
		incReqParams := ctx.Query("id")
		err := json.Unmarshal([]byte(incReqParams), &userID)
		if err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		mappedPictures, err := c.pictureService.GetUserImages(userID)
		if err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		ctx.JSON(http.StatusOK, mappedPictures)
	}
}

func (c *PictureController) GetPictureByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var id int
		incRequestParams := ctx.Query("id")
		err := json.Unmarshal([]byte(incRequestParams), &id)
		if err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		mappedPicture, err := c.pictureService.GetImageByID(id)
		if err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
		}
		ctx.JSON(http.StatusOK, mappedPicture)
	}
}

func (c *PictureController) UpdateImage() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var updateImageRequest models.PictureUpdateRequest
		ctx.ShouldBind(&updateImageRequest)
		identifier, err := strconv.Atoi(ctx.GetHeader("Identifier"))
		if err != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}

		err = c.pictureService.UpdateImageInformation(updateImageRequest.ID, updateImageRequest.Description, updateImageRequest.Hashtags, identifier)
		if err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		ctx.JSON(http.StatusOK, nil)
	}
}

func (c *PictureController) GetEditedImage() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var editImageRequests models.EditImageRequest
		ctx.ShouldBind(&editImageRequests)
		image, err := c.pictureService.GetEditedImage(editImageRequests.ID, editImageRequests.Height, editImageRequests.Width, editImageRequests.Format, editImageRequests.Sepia, editImageRequests.Blur)
		if err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		ctx.Data(http.StatusOK, "application/octet-stream", image.Bytes())
	}
}
