package main

import (
	"log"
	"net/http"

	config "devices/lib/config"
	"devices/lib/devices"

	"github.com/gin-gonic/gin"
)

const Version = "0.1.0"

var cnf config.Config = config.LoadConfig()

func Authenticator() gin.HandlerFunc {
	return func(c *gin.Context) {
		headers := c.Request.Header
		TokenHeader, ok := headers["Token"]

		// If there is not Token header
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": "UnprocessableEntity"})
		}

		// If the token exists and it is correct
		if TokenHeader[0] == cnf.Token {
			return
		}

		// If the token exists bit it is wrong
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "Access Denied"})
	}
}

func main() {
	r := setupRouter()
	r.Run()
}

func setupRouter() *gin.Engine {

	r := gin.Default()

	r.Use(Authenticator())

	r.POST("devices/", func(c *gin.Context) {
		var device devices.Device
		if err := c.ShouldBindJSON(&device); err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Could not parse device status"})
		}

		err := device.AutoTimestamp().Save()

		if err != nil {
			log.Fatal(err)
		}

		c.JSON(http.StatusAccepted, gin.H{"device": device.Name, "status": device.EventMessage()})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	return r
}
