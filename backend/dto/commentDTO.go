package dto

import "gin_app/models"

type commentCreateRequestDTO struct {
	Content string `json:"comment" binding:"required" gorm:"not null"`
}

func CommentCreateRequestDTO() *commentCreateRequestDTO {
	return &commentCreateRequestDTO{}
}

type commentCreateResponseDTO struct {
	ID      uint   `json:"id"`
	UserID  uint   `json:"user_id"`
	PostID  uint   `json:"post_id"`
	Content string `json:"comment"`
}

func CommentCreateResponseDTO(comment models.Comment) *commentCreateResponseDTO {
	return &commentCreateResponseDTO{
		ID:      comment.ID,
		UserID:  comment.UserID,
		PostID:  comment.PostID,
		Content: comment.Content,
	}
}
