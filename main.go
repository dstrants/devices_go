package main

import (
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"

	config "devices/lib/config"
	"devices/lib/devices"

	"github.com/gin-gonic/gin"
)

const Version = "0.1.1"

var cnf config.Config = config.LoadConfig()

// Removes line endings from user input field
func escapeInputField(field string) string {
	escapedField := strings.Replace(field, "\n", "", -1)
	escapedField = strings.Replace(escapedField, "\r", "", -1)

	return escapedField
}

func Authenticator() gin.HandlerFunc {
	return func(c *gin.Context) {
		headers := c.Request.Header
		TokenHeader, ok := headers["Token"]

		// If there is not Token header
		if !ok {
			log.Error("Request did not include a token header")
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": "UnprocessableEntity"})
		}

		// If the token exists and it is correct
		if TokenHeader[0] == cnf.Token {
			return
		}

		// If the token exists bit it is wrong
		log.Error("Authentication failed")
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "Access Denied"})
	}
}

func main() {
	if gin.Mode() == "release" {
		log.SetFormatter(&log.JSONFormatter{})
		log.Info("Logging set to json")
	}

	r := setupRouter()
	r.Run()
}

func setupRouter() *gin.Engine {

	r := gin.Default()

	r.Use(Authenticator())

	r.POST("devices/", func(c *gin.Context) {
		var device devices.Device
		if err := c.ShouldBindJSON(&device); err != nil {
			log.Error("Could not parse request body")
			c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Could not parse device status"})
		}

		// Send slack notification and save the data to mongo
		err := device.Notify()

		if err != nil {
			log.WithError(err).Error("An error occurred while sending notification")
		}

		err = device.AutoTimestamp().Save()

		if err != nil {
			log.WithError(err).WithFields(log.Fields{"device": escapeInputField(device.Name)}).Error("An error occurred while saving event to mongo")
		}

		c.JSON(http.StatusAccepted, gin.H{"device": device.Name, "status": device.EventMessage()})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	return r
}
