package models

import (
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Picture struct {
	gorm.Model
	Title          string
	Description    string
	PictureURI     string
	UploadDateTime time.Time
	Hashtags       pq.StringArray `gorm:"type:text[]"`
	User           User           `gorm:"embedded;embeddedPrefix:user_"`
}
