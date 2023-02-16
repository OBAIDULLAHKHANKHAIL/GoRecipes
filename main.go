package main

import (
	"github.com/gin-gonic/gin"
)

//main function
func main() {
	router := gin.Default()
	router.GET("/home", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})
	router.Run(":5001")
}
