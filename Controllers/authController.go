package controllers

import (
	"gin_app/config"
	"gin_app/dto"
	"gin_app/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	var user models.User
	var existingUser models.User

	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Invalid JSON data",
		})
		return
	}

	result := config.DB.Where("email = ?", user.Email).First(&existingUser)

	if result.RowsAffected > 0 {
		c.JSON(http.StatusConflict, gin.H{
			"status":  http.StatusConflict,
			"message": "Email already registered",
		})
		return
	}

	result = config.DB.Where("username = ?", user.Username).First(&existingUser)
	if result.RowsAffected > 0 {
		c.JSON(http.StatusConflict, gin.H{
			"status":  http.StatusConflict,
			"message": "User already exists",
		})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Couldn't hash the password",
		})
		return
	}
	user.Password = string(hashedPassword)

	err = config.DB.Create(&user).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Couldn't create user",
		})
		return
	}

	userDTO := dto.UserAuthSignUpResponseDTO(user)

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "User created successfully",
		"user":    userDTO,
	})
}

func Login(c *gin.Context) {

	var loginUser = dto.UserAuthLoginRequestDTO()

	var existingUser models.User
	err := c.ShouldBindJSON(&loginUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Invalid Json Data",
		})
		return
	}

	result := config.DB.Where("email = ? OR username = ?", loginUser.LoginField, loginUser.LoginField).Find(&existingUser)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Wrong Email/Username Or Password",
		})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(loginUser.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Wrong Email/Username Or Password",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Successfully Logged In",
		"user": gin.H{
			"ID":         existingUser.ID,
			"LoginField": loginUser.LoginField,
		},
	})

}
