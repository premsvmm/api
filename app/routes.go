package app

import (
	"github.com/gin-gonic/gin"
	"github.com/premsvmm/goapi/controllers"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func routes() {
	Router.GET("", controllers.Welcome)
	Router.GET("/hostname", controllers.Hostname)
	Router.POST("/add",controllers.Addition)
	Router.GET("/metrics", prometheusHandler())
}

func prometheusHandler() gin.HandlerFunc {
	h := promhttp.Handler()

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
func webRoutes(){
	Router.GET("/home",controllers.Webview)
}
