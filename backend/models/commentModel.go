package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	UserID    uint   `json:"user_id" binding:"required" gorm:"not null"`
	PostID    uint   `json:"post_id" binding:"required" gorm:"not null"`
	Content   string `json:"comment" binding:"required" gorm:"not null"`
	UpVotes   uint   `json:"likes"`
	DownVotes uint   `json:"dislikes"`
}
