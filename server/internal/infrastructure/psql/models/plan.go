package models

import "gorm.io/gorm"

type Plan struct {
	gorm.Model
	PlanName          string
	UploadLimitSizeKb uint32
	DailyUploadLimit  uint32
}
