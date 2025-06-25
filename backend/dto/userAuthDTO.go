package dto

import models "gin_app/models"

type UserSignUpRequestDTO struct {
	FirstName     string `json:"first_name" binding:"required"`
	LastName     string `json:"last_name" binding:"required"`
	DateOfBirth string `json:"dob" binding:"required"`
	Gender string `json:"gender" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	CountryCode string `json:"country" binding:"required"`
}



type userSignUpResponseDTO struct {
	ID    uint   `json:"id"`
	FullName  string `json:"full_name"`

}

func UserSignUpResponseDTO(user models.User) *userSignUpResponseDTO {
	return &userSignUpResponseDTO{
		user.ID,
		user.FirstName + " " + user.LastName,
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
	FullName string `json:"full_name"`
}

func UserLoginResponseDTO(id uint, FullName string) *userLoginResponseDTO {
	return &userLoginResponseDTO{
		ID:         id,
		FullName: FullName,
	}
}
