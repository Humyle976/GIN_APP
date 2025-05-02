package dto

import models "gin_app/models"

type userSignUpRequestDTO struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Age      uint8  `json:"age" binding:"required"`
}

func UserSignUpRequestDTO() *userSignUpRequestDTO {
	return &userSignUpRequestDTO{}
}

type userSignUpResponseDTO struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   uint8  `json:"age"`
}

func UserSignUpResponseDTO(user models.User) *userSignUpResponseDTO {
	return &userSignUpResponseDTO{
		user.ID,
		user.Username,
		user.Email,
		user.Age,
	}
}

type userLoginRequestDTO struct {
	LoginField string `json:"loginfield" binding:"required"`
	Password   string `json:"password" binding:"required,min=6"`
}

func UserLoginRequestDTO() *userLoginRequestDTO {
	return &userLoginRequestDTO{}
}

type userLoginResponseDTO struct {
	ID         uint   `json:"user_id"`
	LoginField string `json:"Email/Username"`
}

func UserLoginResponseDTO(id uint, field string) *userLoginResponseDTO {
	return &userLoginResponseDTO{
		ID:         id,
		LoginField: field,
	}
}
