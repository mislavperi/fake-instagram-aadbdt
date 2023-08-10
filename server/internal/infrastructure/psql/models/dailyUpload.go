package models

import "gorm.io/gorm"

type DailyUpload struct {
	gorm.Model
	UploadSizeKb uint64
	Picture      Picture `gorm:"embedded;embeddedPrefix:picture_"`
	User         User    `gorm:"embedded;embeddedPrefix:user_"`
}
