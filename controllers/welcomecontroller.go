package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Welcome(c *gin.Context){
	REQUEST_COUNT.Inc()
	c.String(http.StatusOK,"welcome Prem")
}