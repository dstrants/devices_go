package main

import (
	"net/http"

	"devices/lib/devices"

	"github.com/gin-gonic/gin"
)

func main() {
	r := setupRouter()
	r.Run()
}

func setupRouter() *gin.Engine {

	r := gin.Default()

	r.POST("devices/", func(c *gin.Context) {
		var device devices.Device
		if err := c.ShouldBindJSON(&device); err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Could not parse device status"})
		}
		c.JSON(http.StatusAccepted, gin.H{"device": device.Name, "status": device.EventMessage()})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	return r
}
