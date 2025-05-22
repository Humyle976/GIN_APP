package helpers

import (
	"gin_app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserFromContext(c *gin.Context) (models.User, bool) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "You are not authorized to do that",
		})
		return models.User{}, false
	}
	return user.(models.User), true
}
