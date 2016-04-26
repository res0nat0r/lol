package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("ERROR: $PORT must be set")
	}

	r := gin.Default()

	r.GET("/stefhen", stefhen)
	r.Run(":" + port)
}

func stefhen(c *gin.Context) {
	c.JSON(200, gin.H{"a": "b"})
}
