package models

import "gorm.io/gorm"

type Visibility string

const (
	Public  Visibility = "public"
	Private Visibility = "private"
)

type Post struct {
	gorm.Model
	Visibility Visibility `gorm:"type:visibility_enum;default:'public'" json:"visibility"`
	UserID     uint       `json:"user_id" binding:"required"`
	Content    string     `gorm:"size:100" json:"content" binding:"required"`
	Comments   []Comment  `gorm:"foreignKey:PostID;constraint:OnDelete:CASCADE" json:"-"`
	Likes      []Likes    `gorm:"foreignKey:PostID;constraint:OnDelete:CASCADE" json:"-"`
	FileURL  string `json:"file_url,omitempty"`
}
