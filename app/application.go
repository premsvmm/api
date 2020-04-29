package app

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	Router = gin.Default()
)

func StartApplication() {
	Router.Use(cors.Default())
	routes()
	webRoutes()
	Router.LoadHTMLGlob("templates/*")
	Router.Run(":8080")
}

func StartTestApplication() {
	gin.SetMode(gin.TestMode)
	routes()
}
