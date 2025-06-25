package dto

import models "gin_app/models"

type postCreateRequestDTO struct {
	Content string `binding:"required"`
	Visibility models.Visibility `default:"public"`
}

func PostCreateRequestDTO() *postCreateRequestDTO {
	return &postCreateRequestDTO{}
}



type postGetResponseDTO struct {
	ID       uint   `json:"id"`
	UserID   uint   `json:"user_id" binding:"required"`
	Fullname string `json:"fullname"`
	Content  string `json:"content" binding:"required"`
}

func PostGetResponseDTO(posts []models.Post, user UserContext) []postGetResponseDTO {
	allposts := []postGetResponseDTO{}
	for _, v := range posts {
		allposts = append(allposts, postGetResponseDTO{
			ID:       v.ID,
			UserID:   user.ID,
			Fullname: user.Fullname,
			Content:  v.Content,
		})
	}
	return allposts
}
