package router

import (
	"fmt"
	propertyController "microservicio/controllers/property"
	userController "microservicio/controllers/user"

	"github.com/gin-gonic/gin"
)

func MapUrls(router *gin.Engine) {
	// Products Mapping
	router.GET("/properties/:parameters/id", propertyController.Get)

	router.GET("/properties/:parameters", propertyController.GetByParam)
	router.GET("/properties/all", propertyController.GetAll)
	router.POST("/properties/load", propertyController.Insert)
	router.POST("/properties/import", propertyController.InsertMany)

	router.GET("/user/:id", userController.GetUserById)
	router.POST("/login", userController.UserLogin)

	fmt.Println("Finishing mappings configurations")
}
