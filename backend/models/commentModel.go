package models

type Comment struct {
	ID        uint   `gorm:"primaryKey;not null" json:"id" binding:"required"`
	UserID    uint   `json:"user_id" binding:"required" gorm:"not null"`
	PostID    uint   `json:"post_id" binding:"required" gorm:"not null"`
	Content   string `json:"comment" binding:"required" gorm:"not null"`
	UpVotes   uint   `json:"likes"`
	DownVotes uint   `json:"dislikes"`
}
