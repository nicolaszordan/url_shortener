package main

import (
	"github.com/gin-gonic/gin"
)

// Define the routes for the URL shortener service
func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/shorten", func(c *gin.Context) {
		// TODO: Handle URL shortening
	})
	r.GET("/expand", func(c *gin.Context) {
		// TODO: Handle URL expansion
	})
	return r
}

func main() {
	r := setupRouter()

	r.Run(":8080")
}
