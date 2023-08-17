package models

import (
	"time"

	"gorm.io/gorm"
)

type DailyUpload struct {
	gorm.Model
	ID           int64 `gorm:"primary_key"`
	UploadSizeKb uint64
	PictureID    int64
	Picture      Picture
	UserID       int64
	User         User
	CreatedAt    time.Time
}
