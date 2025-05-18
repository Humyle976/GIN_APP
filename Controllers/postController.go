package controllers

import (
	"gin_app/config"
	"gin_app/dto"
	"gin_app/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeleteAPost(c *gin.Context) {

	postIdStr := c.Param("postId")

	postIDUint64, err := strconv.ParseUint(postIdStr, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Invalid post ID",
		})
		return
	}

	currentPost := uint(postIDUint64)

	post := models.Post{}
	err = config.DB.First(&post, currentPost).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"status":  http.StatusNotFound,
				"message": "Post not found",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Internal Server Error",
			})
		}
		return
	}

	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "You are not authorized to do that",
		})
		return
	}

	currentUser := user.(models.User)

	if post.UserID != currentUser.ID {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "You are not authorized to do that",
		})
		return
	}

	if err := config.DB.Delete(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Failed to delete post",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Post deleted successfully",
	})
}

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
