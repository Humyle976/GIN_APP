package main

import (
	models "gin_app/Models"
	routers "gin_app/Routers"
	"gin_app/config"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	config.ConnectDB()
	config.DB.AutoMigrate(&models.User{})
}
func main() {

	router := gin.Default()
	err := godotenv.Load(".env")

	if err != nil {
		log.Println("No .env file found")
	}

	routers.UserRouter(router)

	router.Run(":8000")

}
