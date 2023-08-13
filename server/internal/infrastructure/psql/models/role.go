package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	ID int64 `gorm:"primary_key"`
	Name string
}
