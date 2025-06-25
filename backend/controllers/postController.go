package controllers

import (
	"fmt"
	"gin_app/config"
	"gin_app/dto"
	"gin_app/helpers"
	"gin_app/models"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

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

	currentUser, ok := helpers.GetUserFromContext(c)
	if !ok {
		return
	}
	
	if post.UserID != currentUser.ID {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "You are not authorized to do that",
		})
		return
	}


	relativePath := strings.TrimLeft(post.FileURL, "/");
	cleanPath := filepath.Clean(relativePath)
	_ = os.Remove(cleanPath)
	
	err = config.DB.Unscoped().Delete(&post).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Failed to delete post",
		})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{
		"status":  http.StatusNoContent,
	})
}

func CreateAPost(c *gin.Context) {
	user, ok := helpers.GetUserFromContext(c)
	if !ok {
		return
	}

	text := c.PostForm("text")
	if text == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Post must contain some text",
		})
		return
	}
	if len(text) > 400 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Text must be less than 400 characters",
		})
		return
	}

	visibility := c.DefaultPostForm("visibility", "public")
	var fileURL string

	file, header, err := c.Request.FormFile("file")
	if err == nil {
		defer file.Close()

		ext := strings.ToLower(filepath.Ext(header.Filename))
		allowed := map[string]bool{
			".jpg":  true,
			".jpeg": true,
			".png":  true,
			".gif":  true,
			".webp": true,
			".mp4":  true,
			".mov":  true,
			".avi":  true,
			".webm": true,
		}
		if !allowed[ext] {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"message": "Unsupported file type",
			})
			return
		}

		fileName := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
		savePath := filepath.Join("uploads", fileName)

		err := c.SaveUploadedFile(header, savePath)
		
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Internal Server Error",
			})
			return
		}

		fileURL = "/uploads/" + fileName

	} else if err != http.ErrMissingFile {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Error reading file",
		})
		return
	}

	post := models.Post{
		UserID:     user.ID,
		Content:    text,
		FileURL:    fileURL,
		Visibility: models.Visibility(visibility),
	}

	if err := config.DB.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Internal Server Error",
		})
		return
	}
	postObj := dto.PostWithUser{
		PostID: post.ID,
		Content: post.Content,
		UserID: post.UserID,
		Fullname: user.Fullname,
		FileURL: post.FileURL,
		CreatedAt: post.CreatedAt,
		Likes: 0,
		Comments: 0,
		IsOwner: true,

	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"post":    postObj,
	})
}


func GetPostsOfCurrentUser(c *gin.Context) {
	user, ok := helpers.GetUserFromContext(c)
	if !ok {
		return
	}

	var posts []models.Post
	err := config.DB.Where("user_id = ?", user.ID).Find(&posts).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Couldn't get user's posts",
		})
		return
	}
	
	postGetResponseDTO := dto.PostGetResponseDTO(posts, user)
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   postGetResponseDTO,
	})

}
