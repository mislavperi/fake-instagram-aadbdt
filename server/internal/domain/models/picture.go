package models

import (
	"mime/multipart"
	"time"
)

type Picture struct {
	ID             int64     `json:"id"`
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	PictureURI     string    `json:"pictureURI"`
	UploadDateTime time.Time `json:"uploadDateTime"`
	Hashtags       []string  `json:"hashtags"`
	User           User      `json:"user"`
}

type PictureUpload struct {
	Title       string                `form:"title" binding:"required"`
	Description string                `form:"description" binding:"required"`
	Picture     *multipart.FileHeader `form:"file" binding:"required"`
	Hashtags    string                `form:"hashtags" binding:"required"`
	Format      string                `form:"format" binding:"required"`
	Height      string                `form:"height" binding:"required"`
	Width       string                `form:"width" binding:"required"`
}

type GetPictureRequest struct {
	ID int `json:"id"`
}

type PictureUpdateRequest struct {
	ID          int      `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Hashtags    []string `json:"hashtags"`
}

type EditImageRequest struct {
	ID     int     `json:"id"`
	Height int32   `json:"height"`
	Width  int32   `json:"width"`
	Format string  `json:"format"`
	Sepia  float32 `json:"sepia"`
	Blur   float32 `json:"blur"`
}
