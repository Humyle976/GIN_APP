package main

import (
	"gin_app/config"
	"gin_app/models"
	"gin_app/routers"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnv()
	config.ConnectDB()
	config.DB.AutoMigrate(&models.User{})
	config.InitRedisClient()
}
func main() {

	router := gin.Default()

	routers.UserRouter(router)
	routers.AuthRouter(router)
	router.Run(os.Getenv("Address") + ":" + os.Getenv("PORT"))

}
