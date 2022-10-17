package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	bookController "microservicio/controllers/book"
)

func MapUrls(router *gin.Engine) {
	// Products Mapping
	router.GET("/properties/:id", bookController.Get)
	router.POST("/properties/load", bookController.Insert)

	fmt.Println("Finishing mappings configurations")
}
