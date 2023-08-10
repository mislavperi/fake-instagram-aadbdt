package controllers

import (
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mislavperi/fake-instagram-aadbdt/server/internal/domain/models"
)

type PictureController struct {
	pictureService PictureService
}

type PictureService interface {
	UploadImage(file multipart.File, title string, description string, hashtags []string, user models.User) error
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

		file, err := pictureUpload.Picture.Open()
		if err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
		}
		c.pictureService.UploadImage(file, pictureUpload.Title, pictureUpload.Description, pictureUpload.Hashtags, pictureUpload.User)
	}
}
