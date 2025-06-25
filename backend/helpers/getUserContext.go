package helpers

import (
	"gin_app/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserFromContext(c *gin.Context) (dto.UserContext, bool) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "You are not authorized to do that",
		})
		return dto.UserContext{}, false
	}
	return user.(dto.UserContext), true
}
