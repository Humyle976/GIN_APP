package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.Engine) {
	user := router.Group("/user")
	{
		user.GET("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusAccepted, gin.H{
				"message": "user",
			})
		})
	}
}
