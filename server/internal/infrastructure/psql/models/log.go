package models

import (
	"time"

	"gorm.io/gorm"
)

type Log struct {
	gorm.Model
	User          User `gorm:"embedded;embeddedPrefix:user_"`
	Action        string
	Timestamp     time.Time
	PictureObject Picture `gorm:"embedded;embeddedPrefix:picture_"`
}
