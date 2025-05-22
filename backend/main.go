package main

import (
	"gin_app/config"
	"gin_app/routers"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnv()
	config.ConnectPostgres()
	config.Migrate()
	config.InitRedisClient()
	config.ConnectTigerGraph()
}
func main() {

	router := gin.Default()

	routers.PostRouter(router)
	routers.UserRouter(router)
	routers.AuthRouter(router)
	router.Run(os.Getenv("Address") + ":" + os.Getenv("PORT"))

}
