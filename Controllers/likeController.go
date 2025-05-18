package controllers

import (
	"gin_app/config"
	"gin_app/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func LikeAPost(c *gin.Context) {

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
	like := &models.Likes{
		PostID:   post.ID,
		UserID:   currentUser.ID,
		Username: currentUser.Username,
	}
	err = config.DB.Model(&models.Likes{}).Create(&like).Error

	if err != nil {
		if err == gorm.ErrDuplicatedKey {
			c.JSON(http.StatusConflict, gin.H{
				"status":  http.StatusConflict,
				"message": "You have already liked the post",
			})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Couldn't like the post",
			})
		}
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Liked the Post",
	})
}

func DislikeAPost(c *gin.Context) {
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
	likes := &models.Likes{}
	err = config.DB.Select("post_id", "user_id").Where("post_id = ? AND user_id = ?", post.ID, currentUser.ID).First(&likes).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"status":  http.StatusNotFound,
				"message": "You have not liked the post",
			})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Internal Server Error",
			})
		}
		return
	}
	err = config.DB.Delete(&likes).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"status":  http.StatusNotFound,
				"message": "You have not liked the post",
			})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Internal Server Error",
			})
		}
		return
	}
	c.JSON(http.StatusNoContent, gin.H{
		"status":  http.StatusNoContent,
		"message": "deleted",
	})

}
func GetAllLikes(c *gin.Context) {
	_, exists := c.Get("user")

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

	likes := []models.Likes{}
	err = config.DB.Where("post_id = ?", post.ID).Find(&likes).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   likes,
	})
}
