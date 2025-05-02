package models

type Post struct {
	ID      uint   `gorm:"primaryKey" json:"id"`
	UserID  uint   `json:"user_id" binding:"required" gorm:"not null"`
	User    User   `json:"-"`
	Content string `gorm:"size:100" json:"content" binding:"required"`
}
