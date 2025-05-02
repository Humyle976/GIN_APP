package routers

import (
	"gin_app/controllers"

	"github.com/gin-gonic/gin"
)

func PostRouter(router *gin.Engine) {
	posts := router.Group("/posts")
	{
		posts.POST("/", controllers.Authenticate, controllers.CreateAPost)
		posts.GET("/", controllers.Authenticate, controllers.GetPostsOfCurrentUser)
	}
}
