package controllers

import (
	"counter-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateStatus(c *gin.Context) {
	status := services.UpdateStatus()

	c.JSON(http.StatusOK, status)
}

func Shutdown(c *gin.Context) {
	services.Shutdown()

	c.JSON(http.StatusOK, services.GetStatus())
}

func PowerOn(c *gin.Context) {
	services.PowerOn()

	c.JSON(http.StatusOK, services.GetStatus())
}
