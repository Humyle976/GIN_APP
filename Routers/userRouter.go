package routers

import (
	"gin_app/controllers"

	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.Engine) {
	users := router.Group("/users")
	{
		users.GET("/", controllers.Authenticate, controllers.GetAllUsers)
	}
}
