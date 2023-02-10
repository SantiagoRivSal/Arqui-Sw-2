package solrController

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"search/services"
	client "search/services/repositories"
	con "search/utils/solr"
)

var (
	Solr = services.NewSolrServiceImpl(
		(*client.SolrClient)(con.NewSolrClient("localhost", 8983, "properties")),
	)
)

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
