package app

import (
	"github.com/premsvmm/goapi/controllers"
)

func routes() {
	Router.GET("", controllers.Welcome)
	Router.GET("/hostname", controllers.Hostname)
}

func webRoutes(){
	Router.GET("/home",controllers.Webview)
}
