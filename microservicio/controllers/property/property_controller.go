package property

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"microservicio/dtos"
	service "microservicio/services"
	"microservicio/utils/cache"
	"net/http"
)

func Get(c *gin.Context) {

	id := c.Param("id")

	res := cache.Get(id)

	if res != "" {
		var propertyDtoCache dtos.PropertyDto
		json.Unmarshal([]byte(res), &propertyDtoCache)
		fmt.Println("from cache: " + id)
		c.JSON(http.StatusOK, propertyDtoCache)
		return

	}

	propertyDto, er := service.PropertyService.GetProperty(id)

	// Error del Insert
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusOK, propertyDto)
}

/*
if (param == "city") {
		propertiesDto, er := service.PropertyService.GetByParam("city")
	} else if (param == "country") {
		propertiesDto, er := service.PropertyService.GetByParam("country")
	} else if (param == "service") {
		propertiesDto, er := service.PropertyService.GetByParam("service")
	}
*/

func GetByParam(c *gin.Context) {
	param := c.Param("param")

	propertiesDto, er := service.PropertyService.GetByParam(param)
	//error del get

	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusOK, propertiesDto)
}
func GetAll(c *gin.Context) {

	propertiesDto, er := service.PropertyService.GetProperties()
	//error del get
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusOK, propertiesDto)
}

func InsertMany(c *gin.Context) {
	var propertiesDto dtos.PropertiesDto
	err := c.BindJSON(&propertiesDto)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	propertiesDto, er := service.PropertyService.InsertMany(propertiesDto)

	if er != nil {
		c.JSON(er.Status(), er)
		return
	}
	c.JSON(http.StatusCreated, propertiesDto)
}

func Insert(c *gin.Context) {
	var propertyDto dtos.PropertyDto
	err := c.BindJSON(&propertyDto)

	// Error Parsing json param
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	propertyDto, er := service.PropertyService.InsertProperty(propertyDto)

	// Error del Insert
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	propertyDtoStr, _ := json.Marshal(propertyDto)

	cache.Set(propertyDto.Id, propertyDtoStr)

	fmt.Println("save cache: " + propertyDto.Id)
	c.JSON(http.StatusCreated, propertyDto)
}

/*
func GetService(c *gin.Context) {

	propertiesDto, er := service.PropertyService.GetServices()
	//error del get
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusOK, propertiesDto)
}

func GetCounty(c *gin.Context) {

	propertiesDto, er := service.PropertyService.GetCountries()
	//error del get
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusOK, propertiesDto)
}

//func GetRandom(c *gin.Context) {}
*/
