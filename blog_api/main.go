package main

import (
	"blog_api/config"
	"blog_api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()
	gin.SetMode(gin.ReleaseMode)
	r := routes.SetupRoutes()
	r.Run(":8080") // Run on port 8080
}
