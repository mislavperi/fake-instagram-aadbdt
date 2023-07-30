package models

import "gorm.io/gorm"

type Package struct {
	gorm.Model
	PackageName       string
	UploadLimitSizeKb uint32
	DailyUploadLimit  uint32
	MaximumStorageKb  uint32
}
