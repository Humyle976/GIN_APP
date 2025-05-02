package controllers

import (
	"gin_app/config"
	"gin_app/dto"
	"gin_app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateAPost(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "You are not authorized to do that",
		})
		return
	}

	postCreateRequestDTO := dto.PostCreateRequestDTO()

	err := c.ShouldBindJSON(postCreateRequestDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Invalid JSON data",
		})
		return
	}
	currentUser := user.(models.User)
	post := models.Post{
		UserID:  currentUser.ID,
		Content: postCreateRequestDTO.Content,
	}

	err = config.DB.Create(&post).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Couldn't create post",
		})
		return
	}

	postCreateResponseDTO := dto.PostCreateResponseDTO(currentUser, post)

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Post created",
		"data":    postCreateResponseDTO,
	})
}

func GetPostsOfCurrentUser(c *gin.Context) {
	user, exists := c.Get("user")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "You are not authorized to do that",
		})
		return
	}

	currentUser := user.(models.User)

	var posts []models.Post
	err := config.DB.Where("user_id = ?", currentUser.ID).Find(&posts).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Couldn't get user's posts",
		})
		return
	}

	postGetResponseDTO := dto.PostGetResponseDTO(posts, currentUser)
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   postGetResponseDTO,
	})

}
