package main

import (
	config "gin_app/Configs"
	models "gin_app/Models"
	routers "gin_app/Routers"

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

	router.Run(":8000")

}
