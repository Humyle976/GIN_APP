package dto

import models "gin_app/models"

type postCreateRequestDTO struct {
	Content string `json:"content" binding:"required"`
}

func PostCreateRequestDTO() *postCreateRequestDTO {
	return &postCreateRequestDTO{}
}

type postCreateResponseDTO struct {
	ID       uint   `json:"id"`
	UserID   uint   `json:"user_id" binding:"required"`
	Username string `json:"username"`
	Content  string `json:"content" binding:"required"`
}

func PostCreateResponseDTO(user models.User, post models.Post) *postCreateResponseDTO {
	return &postCreateResponseDTO{
		ID:       post.ID,
		UserID:   user.ID,
		Username: user.Username,
		Content:  post.Content,
	}
}

type postGetResponseDTO struct {
	ID       uint   `json:"id"`
	UserID   uint   `json:"user_id" binding:"required"`
	Username string `json:"username"`
	Content  string `json:"content" binding:"required"`
}

func PostGetResponseDTO(posts []models.Post, user models.User) []postGetResponseDTO {
	allposts := []postGetResponseDTO{}
	for _, v := range posts {
		allposts = append(allposts, postGetResponseDTO{
			ID:       v.ID,
			UserID:   user.ID,
			Username: user.Username,
			Content:  v.Content,
		})
	}
	return allposts
}
