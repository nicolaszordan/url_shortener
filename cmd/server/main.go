package main

import (
	"github.com/gin-gonic/gin"
)

// Define the routes for the URL shortener service
func setupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/shorten", func(c *gin.Context) {
		// TODO: Handle URL shortening
	})
	r.GET("/shorten/:shortened_url", func(c *gin.Context) {
		// TODO: Handle URL expansion
	})
	r.PUT("/shorten/:shortened_url", func(c *gin.Context) {
		// TODO: Handle URL update
	})
	r.DELETE("/shorten/:shortened_url", func(c *gin.Context) {
		// TODO: Handle URL deletion
	})
	return r
}

func main() {
	r := setupRouter()

	r.Run(":8080")
}
