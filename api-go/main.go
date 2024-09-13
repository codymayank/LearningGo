package main

import (
	"api-go/albums"
	"github.com/gin-gonic/gin" // Importing gin package
	"log"
	"net/http"
)

// ValidateToken is a middleware function that validates the token
func ValidateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		route := c.FullPath()
		if route == "/health" {
			c.Next()
			return
		}
		token := c.GetHeader("Authorization")
		if token != "test" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}
		c.Next()
	}
}

func main() {
	router := gin.Default()
	// request router to use ValidateToken middleware
	router.Use(ValidateToken())
	albums.RegisterAlbumRouters(router)
	err := router.Run("localhost:8080")
	if err != nil {
		log.Fatalf("Error while starting the server: %v", err)
	}
}
