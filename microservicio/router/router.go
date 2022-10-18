package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	propertyController "microservicio/controllers/property"
)

func MapUrls(router *gin.Engine) {
	// Products Mapping
	router.GET("/properties/:id", propertyController.Get)
	router.POST("/properties/load", propertyController.Insert)

	fmt.Println("Finishing mappings configurations")
}
