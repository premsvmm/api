package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/premsvmm/goapi/model/hostname"
	"net/http"
)

func Hostname(c *gin.Context) {
	res := hostname.GetHostname()
	c.JSON(http.StatusOK, res)
}


