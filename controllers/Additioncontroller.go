package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/premsvmm/goapi/model/addition"
	"net/http"
)

func Addition(c *gin.Context) {
	REQUEST_COUNT.Inc()
	var a addition.Add
	c.ShouldBindJSON(&a)
	c.JSON(http.StatusOK, addition.Addition(a))
}
