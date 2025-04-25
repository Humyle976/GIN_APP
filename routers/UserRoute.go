package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.Engine) {

	users := router.Group("/users")
	{
		users.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusAccepted, gin.H{
				"message": "users",
			})
		})

		users.POST("/", func(c *gin.Context) {
			c.JSON(http.StatusCreated, gin.H{
				"message": "User Created!",
			})
		})

		users.GET("/:id", func(c *gin.Context) {
			id := c.Param("id")

			c.JSON(http.StatusCreated, gin.H{
				"message": id + " Created",
			})
		})
	}

}
