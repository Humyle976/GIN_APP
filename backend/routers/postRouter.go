package routers

import (
	"gin_app/controllers"

	"github.com/gin-gonic/gin"
)

func PostRouter(router *gin.Engine) {
	posts := router.Group("/posts")
	posts.Use(controllers.Authorization)
	{
		posts.POST("/", controllers.CreateAPost)
		posts.DELETE("/:postId", controllers.DeleteAPost)

		comments := posts.Group("/:postId/comments")
		{
			comments.POST("/", controllers.AddAComment)
			comments.DELETE("/:commentId", controllers.DeleteAComment)
		}
		likes := posts.Group("/:postId/likes")
		{
			likes.POST("/", controllers.LikeAPost)
			likes.DELETE("/", controllers.DislikeAPost)
			likes.GET("/", controllers.GetAllLikes)
		}
	}
}
