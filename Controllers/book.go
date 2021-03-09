package Controllers

import (
	"net/http"

	"github.com/eaarda/go-restapi-sample/models"

	"github.com/gin-gonic/gin"
)

func GetAllBook(c *gin.Context) {
	var book []models.Book

	err := models.Book
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, book)
	}

}
