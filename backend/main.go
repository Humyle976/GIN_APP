package main

import (
	"gin_app/config"
	"gin_app/routers"
	"os"

	"github.com/gin-contrib/cors"
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
	
	router.Use(cors.New(cors.Config{
    AllowOrigins:     []string{"http://localhost:5173" , "http://192.168.1.108:5173"},
    AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE" , "OPTIONS"},
    AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
    ExposeHeaders:    []string{"Content-Length"},
    AllowCredentials: true,
	AllowOriginFunc: func(origin string) bool {
        return origin == "http://localhost:5173" || origin == "http://192.168.1.108:5173"
    },
}))

router.Static("/uploads", "./uploads")
router.OPTIONS("/*path", func(c *gin.Context) {
    c.Status(200)
})

	routers.FeedRouter(router)
	routers.PostRouter(router)
	routers.UserRouter(router)
	routers.AuthRouter(router)
	router.Run(os.Getenv("Address") + ":" + os.Getenv("PORT"))

}
