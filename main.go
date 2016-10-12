package main

import (
	"github.com/gin-gonic/gin"
	// "net/http"
)

func main() {
	router := gin.Default()

	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(200, gin.H{"user": name, "status": "ok"})
	})

	router.Run(":4545")
}
