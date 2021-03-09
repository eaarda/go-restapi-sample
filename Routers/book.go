package Routers

import (
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/test-api")
	{
		grp1.GET("book", Controllers.GetBooks)
	}
}
