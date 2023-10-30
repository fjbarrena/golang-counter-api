package main

import (
	"counter-api/controllers"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	// Allow CORS requests
	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowWildcard:   true,
		AllowMethods:    []string{"*"},
		AllowHeaders:    []string{"*"},
		MaxAge:          86400,
	}))

	port := os.Getenv("GOPORT")

	if port == "" {
		port = ":40000"
	} else {
		port = ":" + port
	}

	r.GET("/status", controllers.UpdateStatus)

	// I will prefer to use a POST here, but let the GET for simplicity (easy to test in the browser :P)
	r.GET("/shutdown", controllers.Shutdown)
	r.GET("/poweron", controllers.PowerOn)

	r.Run(port)
}
