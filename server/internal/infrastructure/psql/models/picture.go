package models

import (
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Picture struct {
	gorm.Model
	ID             int64 `gorm:"primary_key"`
	Title          string
	Description    string
	PictureURI     string
	UploadDateTime time.Time
	Hashtags       pq.StringArray `gorm:"type:text[]"`
	UserID         int64
	User           User
}
