package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// Create a default Gin router
	r := gin.Default()

	// Define a simple GET route
	r.GET("/hello", helloHandler)

	// Start the server on port 8080
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}

// helloHandler handles the /hello route
func helloHandler(c *gin.Context) {
	c.String(200, "Hello, GoCI Lint with Gin!")
}
