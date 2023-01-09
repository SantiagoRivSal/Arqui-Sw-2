package app

import (
	log "github.com/sirupsen/logrus"
	solrController "wesolr/controllers/solr"
)

func mapUrls() {

	router.GET("/search=:searchQuery", solrController.GetQuery)

	router.GET("/searchAll=:searchQuery", solrController.GetQueryAllFields)

	router.GET("/properties/:id", solrController.AddFromId)

	log.Info("Finishing mappings configurations")
}