package dto

import "time"

type PostWithUser struct {
    PostID    uint      `gorm:"column:post_id"`
    Content   string    `gorm:"column:content"`
    UserID    uint      `gorm:"column:user_id"`
    Fullname  string    `gorm:"column:fullname"`
    FileURL   string    `gorm:"column:file_url"`
    CreatedAt time.Time `gorm:"column:created_at"`
    Likes     uint      `gorm:"column:likes"`
    Comments  uint      `gorm:"column:comments"`
    IsOwner   bool      `gorm:"column:is_owner"`
}