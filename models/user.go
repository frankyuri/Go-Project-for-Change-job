package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	DisplayID uint   `gorm:"uniqueIndex"`
	Username  string `json:"username" gorm:"uniqueIndex;not null"`
	Password  string `json:"-" gorm:"not null"`
	Email     string `json:"email" gorm:"uniqueIndex;not null"`
}
