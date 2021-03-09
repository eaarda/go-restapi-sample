package main

import (
	"api-test/Config"
	"api-test/Controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	Config.InitialMigration() //db connect

	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "API Test",
		})
	})

	r.GET("/books", Controllers.GetBooks)
	r.POST("/books", Controllers.CreateBook)

	r.Run(":5000")
}
