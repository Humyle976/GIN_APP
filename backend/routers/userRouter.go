package routers

import (
	"gin_app/controllers"

	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.Engine) {
	users := router.Group("/users")
	users.Use(controllers.Authorization)
	{
		users.GET("/", controllers.GetAllUsers)
		users.GET("/me/posts", controllers.GetPostsOfCurrentUser)
		users.GET("/me/friends", controllers.GetFriendsOfCurrentUser)

		users.POST("/:id/friend-request", controllers.SendFriendRequest)
		users.DELETE("/me/friend-request/:id", controllers.DeclineFriendRequest)
		users.POST("/me/accept-request/:id", controllers.AcceptFriendRequest)
		users.DELETE("/:id/friend-request", controllers.DeleteFriendRequest)
		users.POST("/:id/block", controllers.BlockAUser)
		users.DELETE("/:id/block", controllers.UnblockAUser)

	}
}
