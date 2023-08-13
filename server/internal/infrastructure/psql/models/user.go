package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID        int `gorm:"primary_key"`
	FirstName string
	LastName  string
	Username  string
	Password  string
	Email     string
	RoleID    *int64
	Role      *Role
	PlanID    *int64
	Plan      *Plan
}
