package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string  `gorm:"not null" json:"first_name" binding:"required"`
	LastName string `gorm:"not null" json:"last_name" binding:"required"`
	DateOfBirth time.Time   `gorm:"type:date" json:"dob" binding:"required"`
	CountryCode string `gorm:"size:2" json:"country_code"`
	Gender string `gorm:"type:gender" json:"gender" binding:"required"`
	Email    string  `gorm:"uniqueIndex;not null" json:"email" binding:"required,email"`
	Password string  `gorm:"not null" json:"password" binding:"required,min=6"`
    Country  Country `gorm:"foreignKey:CountryCode;references:Code"`
	Posts    []Post  `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"-"`
	Likes    []Likes `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"-"`
	Visibility Visibility `gorm:"type:visibility_enum;default:'public'" json:"visibility"`
}
