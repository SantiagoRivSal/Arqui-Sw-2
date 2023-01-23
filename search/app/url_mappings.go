package app

import (
	solrController "search/controllers/solr"

	log "github.com/sirupsen/logrus"
)

func mapUrls() {

	router.GET("/search=:searchQuery", solrController.GetQuery)

	router.GET("/searchAll=:searchQuery", solrController.GetQueryAllFields)

	router.GET("/properties/:id", solrController.AddFromId)

	router.DELETE("/items/:id", solrController.Delete)

	log.Info("Finishing mappings configurations")
}
