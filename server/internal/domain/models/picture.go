package models

import (
	"mime/multipart"
	"time"
)

type Picture struct {
	Title          string
	Description    string
	PictureURI     string
	UploadDateTime time.Time
	Hashtags       []string
	User           User
}

type PictureUpload struct {
	Title       string                `form:"title" binding:"required"`
	Description string                `form:"description" binding:"required"`
	Picture     *multipart.FileHeader `form:"file" binding:"required"`
	Hashtags    []string              `form:"hashtags" binding:"required"`
	User        User                  `form:"user" binding:"required"`
}
