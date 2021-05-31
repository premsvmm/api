package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/premsvmm/goapi/model/hostname"
	"net/http"
)

func Webview(c *gin.Context)  {
	REQUEST_COUNT.Inc()
	c.HTML(http.StatusOK,"index.html",gin.H{"hostname":hostname.Get()})
}
