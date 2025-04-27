package main

import (
	routers "gin_app/Routers"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	routers.UserRouter(router)

	router.Run(":8000")

}
