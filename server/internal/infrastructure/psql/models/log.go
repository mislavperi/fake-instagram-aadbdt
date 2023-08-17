package models

import (
	"time"

	"gorm.io/gorm"
)

type Log struct {
	gorm.Model
	ID        int64 `gorm:"primary_key"`
	UserID    int64
	User      User
	Action    string
	Timestamp time.Time
}
