package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/stefhen", stefhen)
	r.Run(":8080")
}

func stefhen(c *gin.Context) {
	c.JSON(200, gin.H{"a": "b"})
}
