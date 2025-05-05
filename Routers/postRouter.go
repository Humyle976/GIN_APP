package routers

import (
	"gin_app/controllers"

	"github.com/gin-gonic/gin"
)

func PostRouter(router *gin.Engine) {
	posts := router.Group("/posts")
	{
		posts.POST("/", controllers.Authenticate, controllers.CreateAPost)
		posts.DELETE("/:postId", controllers.Authenticate, controllers.DeleteAPost)

		comments := posts.Group("/:postId/comments")
		{
			comments.POST("/", controllers.Authenticate, controllers.AddAComment)
			comments.DELETE("/:commentId", controllers.Authenticate, controllers.DeleteAComment)
		}
	}
}
