package app

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	Router = gin.Default()
)

func StartApplication() {
	Router.Use(cors.Default()) // frontend application ot allow cros
	webRoutes()
	Router.LoadHTMLGlob("templates/*")
	routes()
	Router.Run(":8080")
}

func StartTestApplication() {
	gin.SetMode(gin.TestMode)
	routes()
}
