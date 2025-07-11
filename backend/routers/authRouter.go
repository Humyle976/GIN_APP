package routers

import (
	"gin_app/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRouter(router *gin.Engine) {
	auth := router.Group("/auth")
	{
		auth.POST("/signup", controllers.SignUp)
		auth.POST("/verify",controllers.VerifyRegistration)
		auth.GET("/verify", controllers.VerifyTokenUrl)

		auth.POST("/login", controllers.Login)
		auth.GET("/login", controllers.CheckLoginStatus)
		auth.POST("/logout", controllers.Logout)

		auth.GET("/email-exists", controllers.CheckEmailExists)
	}
}
