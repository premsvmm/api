package app

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	Router = gin.Default()
	Prom = gin.New()
)

var
	REQUEST_COUNT = promauto.NewCounter(prometheus.CounterOpts{
		Name: "go_app_requests_count",
		Help: "Total App HTTP Requests count.",
	})

func StartApplication() {
	Router.Use(cors.Default()) // frontend application ot allow cros
	REQUEST_COUNT.Inc()
	webRoutes()
	Router.LoadHTMLGlob("templates/*")
	routes()
	Router.Run(":8080")
}

func StartTestApplication() {
	gin.SetMode(gin.TestMode)
	routes()
}
