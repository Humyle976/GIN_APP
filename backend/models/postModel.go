package models

type Post struct {
	ID       uint      `gorm:"primaryKey" json:"post_id"`
	UserID   uint      `json:"user_id" binding:"required"`
	Content  string    `gorm:"size:100" json:"content" binding:"required"`
	Comments []Comment `gorm:"foreignKey:PostID;constraint:OnDelete:CASCADE" json:"-"`
	Likes    []Likes   `gorm:"foreignKey:PostID;constraint:OnDelete:CASCADE" json:"-"`
}
