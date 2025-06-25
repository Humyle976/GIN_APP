package routers

import (
	"gin_app/controllers"

	"github.com/gin-gonic/gin"
)

func FeedRouter(router *gin.Engine) {
	router.GET("/feed", controllers.Authorization,controllers.GetFeed);
}