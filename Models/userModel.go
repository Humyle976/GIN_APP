package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"uniqueIndex;not null" json:"name" binding:"required"`
	Email    string `gorm:"uniqueIndex;not null" json:"email" binding:"required,email"`
	Password string `gorm:"not null" json:"password" binding:"required,min=6"`
	Age      uint8  `gorm:"not null" json:"age" binding:"required"`
}
