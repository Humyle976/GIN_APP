package dto

import models "gin_app/models"

type userSignupResponseDTO struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   uint8  `json:"age"`
}

type loginRequestDTO struct {
	LoginField string `json:"LoginField" binding:"required"`
	Password   string `json:"password" binding:"required,min=6"`
}

func UserAuthLoginRequestDTO() *loginRequestDTO {
	return &loginRequestDTO{}
}

func UserAuthSignUpResponseDTO(user models.User) *userSignupResponseDTO {
	return &userSignupResponseDTO{
		user.ID,
		user.Username,
		user.Email,
		user.Age,
	}
}
