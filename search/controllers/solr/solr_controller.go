package solrController

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"search/dto"
	"search/services"
	client "search/services/repositories"
	con "search/utils/solr"
)

var (
	Solr = services.NewSolrServiceImpl(
		(*client.SolrClient)(con.NewSolrClient("solr", 8983, "properties")),
	)
)

func GetQuery(c *gin.Context) {
	var propertiesDto dto.PropertiesDto
	query := c.Param("searchQuery")

	propertiesDto, err := Solr.GetQuery(query)
	if err != nil {
		log.Debug(propertiesDto)
		c.JSON(http.StatusBadRequest, propertiesDto)
		return
	}

	c.JSON(http.StatusOK, propertiesDto)

}

func GetQueryAllFields(c *gin.Context) {
	var propertiesDto dto.PropertiesDto
	query := c.Param("searchQuery")

	propertiesDto, err := Solr.GetQueryAllFields(query)
	if err != nil {
		log.Debug(propertiesDto)
		c.JSON(http.StatusBadRequest, propertiesDto)
		return
	}

	c.JSON(http.StatusOK, propertiesDto)

}

func AddFromId(c *gin.Context) {
	id := c.Param("id")
	err := Solr.AddFromId(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusCreated, err)
}

func Delete(c *gin.Context) {
	id := c.Param("id")
	err := Solr.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusCreated, err)
}
