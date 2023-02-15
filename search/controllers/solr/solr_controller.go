package solrController

import (
	"net/http"
	"search/dto"
	"search/services"
	client "search/services/repositories"
	con "search/utils/solr"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

var (
	Solr = services.NewSolrServiceImpl(
		(*client.SolrClient)(con.NewSolrClient("localhost", 8983, "property")),
	)
)

func GetQuery(c *gin.Context) {
	var propertiesDto dto.PropertiesDto
	query := c.Param("solrQuery")

	propertiesDto, err := Solr.GetQuery(query)
	if err != nil {
		log.Debug(propertiesDto)
		c.JSON(http.StatusBadRequest, propertiesDto)
		return
	}

	c.JSON(http.StatusOK, propertiesDto)

}

func Add(c *gin.Context) {
	id := c.Param("id")
	err := Solr.Add(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{})
}

/*
func Delete(c *gin.Context) {
	id := c.Param("id")
	err := Solr.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusCreated, err)
}
*/
