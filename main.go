package main

import (
	"gin_app/config"
	"gin_app/models"
	"gin_app/routers"

	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnv()
	config.ConnectDB()
	config.DB.AutoMigrate(&models.User{})
}
func main() {

	router := gin.Default()

	routers.UserRouter(router)
	routers.AuthRouter(router)
	router.Run(":8000")

}
