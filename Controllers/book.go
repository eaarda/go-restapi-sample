package Controllers

import (
	"api-test/Config"
	"api-test/Models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateBookInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
	Page   int    `json:"page" binding:"required"`
}

func GetBooks(c *gin.Context) {
	var books []Models.Book
	Config.DB.Find(&books)

	c.JSON(http.StatusOK, gin.H{"data": books})
}

func CreateBook(c *gin.Context) {
	var input CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	book := Models.Book{Title: input.Title, Author: input.Author, Page: input.Page}
	Config.DB.Create(&book)
	c.JSON(http.StatusOK, gin.H{"data": book})
}
