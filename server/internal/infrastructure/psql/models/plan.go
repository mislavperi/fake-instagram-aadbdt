package models

import "gorm.io/gorm"

type Plan struct {
	gorm.Model
	ID                int64 `gorm:"primary_key"`
	PlanName          string
	UploadLimitSizeKb uint32
	DailyUploadLimit  uint32
	Cost              uint32
}
