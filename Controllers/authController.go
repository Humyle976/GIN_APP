package controllers

import (
	"context"
	"gin_app/config"
	"gin_app/dto"
	"gin_app/helpers"
	"gin_app/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	userSignUpRequestDTO := dto.UserSignUpRequestDTO()
	var existingUser models.User

	err := c.ShouldBindJSON(userSignUpRequestDTO)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Invalid JSON data",
		})
		return
	}

	result := config.DB.Where("email = ?", userSignUpRequestDTO.Email).First(&existingUser)

	if result.RowsAffected > 0 {
		c.JSON(http.StatusConflict, gin.H{
			"status":  http.StatusConflict,
			"message": "Email already registered",
		})
		return
	}

	result = config.DB.Where("username = ?", userSignUpRequestDTO.Name).First(&existingUser)
	if result.RowsAffected > 0 {
		c.JSON(http.StatusConflict, gin.H{
			"status":  http.StatusConflict,
			"message": "User already exists",
		})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userSignUpRequestDTO.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Couldn't hash the password",
		})
		return
	}

	user := models.User{
		Username: userSignUpRequestDTO.Name,
		Email:    userSignUpRequestDTO.Email,
		Password: string(hashedPassword),
		Age:      userSignUpRequestDTO.Age,
	}

	err = config.DB.Create(&user).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Couldn't create user",
		})
		return
	}

	userSignUpResponseDTO := dto.UserSignUpResponseDTO(user)

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "User created successfully",
		"user":    userSignUpResponseDTO,
	})
}

func Login(c *gin.Context) {

	userLoginRequestDTO := dto.UserLoginRequestDTO()

	var existingUser models.User
	err := c.ShouldBindJSON(userLoginRequestDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Invalid Json Data",
		})
		return
	}

	result := config.DB.Where("email = ? OR username = ?", userLoginRequestDTO.LoginField, userLoginRequestDTO.LoginField).Find(&existingUser)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "Wrong Email/Username Or Password",
		})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(userLoginRequestDTO.Password))

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "Wrong Email/Username Or Password",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": existingUser.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Internal Server Error",
		})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", true, true)

	userLoginResponseDTO := dto.UserLoginResponseDTO(existingUser.ID, userLoginRequestDTO.LoginField)
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Successfully Logged In",
		"user":    userLoginResponseDTO,
	})

}

func CheckLoginStatus(c *gin.Context) {
	tokenString, _ := c.Cookie("Authorization")

	if tokenString != "" {
		claims, err := helpers.VerifyJWT(tokenString)

		if err != nil {
			c.SetCookie("Authorization", "", -1, "", "", true, true)
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  http.StatusUnauthorized,
				"message": "Invalid or expired token",
			})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status":  http.StatusOK,
				"message": "You are already logged in",
				"user": gin.H{
					"ID": claims["sub"],
				},
			})
			return
		}
	}

	c.JSON(http.StatusUnauthorized, gin.H{
		"status":  http.StatusUnauthorized,
		"message": "You are not Logged In",
	})
}

func Logout(c *gin.Context) {

	token, err := c.Cookie("Authorization")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "You are not logged in",
		})
		return
	}

	claims, err := helpers.VerifyJWT(token)

	if err != nil {
		c.SetCookie("Authorization", "", -1, "", "", true, true)

		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "Invalid or expired token",
		})
		return
	}

	ctx := context.Background()
	expTime := claims["exp"].(float64) - float64(time.Now().Unix())

	pipe := config.Client.Pipeline()
	pipe.SetNX(ctx, `auth:blacklist:`+token, "", time.Duration(expTime)*time.Second)
	_, err = pipe.Exec(ctx)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Could not log out",
		})

		return
	}
	c.SetCookie("Authorization", "", -1, "", "", true, true)

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Successfully Logged Out",
	})
}

func Authenticate(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "You are not authorized to do that",
		})
		c.Abort()
		return
	}

	claims, err := helpers.VerifyJWT(tokenString)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "You are not authorized to do that",
		})
		c.Abort()
		return
	}

	var user models.User
	config.DB.Select("id", "username").Find(&user, claims["sub"])

	if user.ID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "You are not authorized to do that",
		})
		c.Abort()
		return
	}

	c.Set("user", user)
	c.Next()

}
