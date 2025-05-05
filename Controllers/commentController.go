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

func AddAComment(c *gin.Context) {

	user, exists := c.Get("user")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "You are not authorized to do that",
		})
		return
	}

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

	err = config.DB.First(&models.Post{}, currentPost).Error
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

	addACommentRequestDTO := dto.CommentCreateRequestDTO()

	err = c.ShouldBindJSON(addACommentRequestDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Invalid JSON data",
		})
		return
	}

	currentUser := user.(models.User)

	comment := models.Comment{
		UserID:  currentUser.ID,
		PostID:  currentPost,
		Content: addACommentRequestDTO.Content,
	}

	err = config.DB.Create(&comment).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Couldn't add a comment",
		})
		return
	}

	addACommentResponseDTO := dto.CommentCreateResponseDTO(comment)

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Added comment",
		"data":    addACommentResponseDTO,
	})
}

func DeleteAComment(c *gin.Context) {

	user, exists := c.Get("user")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "You are not authorized to do that",
		})
		return
	}
	IdStr := c.Param("commentId")

	IDUint64, err := strconv.ParseUint(IdStr, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Invalid post ID",
		})
		return
	}

	currentComment := uint(IDUint64)

	comment := models.Comment{}
	err = config.DB.First(&comment, currentComment).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"status":  http.StatusNotFound,
				"message": "Comment not found",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Internal Server Error",
			})
		}
		return
	}

	IdStr = c.Param("postId")

	IDUint64, err = strconv.ParseUint(IdStr, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Invalid post ID",
		})
		return
	}

	currentPost := uint(IDUint64)

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
				"message": "Internal Server Issue",
			})
		}
		return
	}

	currentUser := user.(models.User)

	if comment.UserID != currentUser.ID && post.UserID != currentUser.ID {
		c.JSON(http.StatusForbidden, gin.H{
			"status":  http.StatusForbidden,
			"message": "You do not have permission to delete this comment",
		})
		return
	}

	if err := config.DB.Delete(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Failed to delete comment",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Comment deleted successfully",
	})
}
