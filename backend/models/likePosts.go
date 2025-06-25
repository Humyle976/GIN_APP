package models

type Likes struct {
	PostID   uint   `gorm:"primaryKey" json:"post_id" binding:"required"`
	UserID   uint   `gorm:"primaryKey" json:"user_id" binding:"required"`
}
