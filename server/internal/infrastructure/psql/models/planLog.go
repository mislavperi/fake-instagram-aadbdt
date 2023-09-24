package models

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type PlanLog struct {
	gorm.Model
	ID          int64 `gorm:"primary_key"`
	PlanID      int64 `gorm:"default:1"`
	Plan        Plan
	UserID      int64
	User        User
	ActivatedAt sql.NullTime
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
