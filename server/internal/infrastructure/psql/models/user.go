package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName string
	LastName  string
	Username  string
	Password  string
	Email     string
	Role      Role `gorm:"embedded;embeddedPrefix:role_"`
	Plan      Plan `gorm:"embedded;embeddedPrefix:plan_"`
}
