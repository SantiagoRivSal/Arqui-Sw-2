package router

import (
	"fmt"
	propertyController "microservicio/controllers/property"

	"github.com/gin-gonic/gin"
)

func MapUrls(router *gin.Engine) {
	// Products Mapping
	router.GET("/properties/:id", propertyController.Get)
	router.GET("/properties/all", propertyController.GetAll)
	router.POST("/properties/load", propertyController.Insert)
	router.POST("/properties/import", propertyController.InsertMany)

	//router.GET("/propertyRandom/:cantidad", propertyController.GetRandom)
	router.GET("/properties/country", propertyController.GetCounty)
	router.GET("/properties/city", propertyController.GetCity)
	router.GET("/properties/service", propertyController.GetService)
	fmt.Println("Finishing mappings configurations")
}
