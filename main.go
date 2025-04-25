package main

import (
	"fmt"
	"gin-learning/routers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusAccepted, gin.H{
			"message": "Hello, Gin!",
		})
	})

	routers.UserRouter(router)

	err := router.Run(":8000")
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
