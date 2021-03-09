package main

import (
	"api-test/Config"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "API Test",
		})
	})
	Config.InitialMigration()
	r.Run(":5000")
}
