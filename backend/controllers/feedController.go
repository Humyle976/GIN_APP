package controllers

import (
	"gin_app/config"
	"gin_app/dto"
	"gin_app/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetFeed(c *gin.Context) {
	user, ok := helpers.GetUserFromContext(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	var results []dto.PostWithUser

	query := `
	SELECT 
	  posts.id AS post_id,
	  users.id AS user_id,
	  users.first_name || ' ' || users.last_name AS fullname,
	  posts.created_at,
	  posts.content,
	  posts.file_url,
	  COUNT(DISTINCT likes.post_id) AS likes,
	  COUNT(DISTINCT comments.id) AS comments,
	  CASE 
	    WHEN posts.user_id = ? THEN true 
	    ELSE false 
	  END AS is_owner
	FROM posts
	JOIN users ON posts.user_id = users.id
	LEFT JOIN likes ON posts.id = likes.post_id
	LEFT JOIN comments ON posts.id = comments.post_id
	GROUP BY 
	  posts.id, 
	  users.id, 
	  users.first_name, 
	  users.last_name, 
	  posts.created_at, 
	  posts.content, 
	  posts.file_url
	ORDER BY posts.created_at DESC;
	`

	err := config.DB.Raw(query, user.ID).Scan(&results).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Database error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   results,
	})
}
