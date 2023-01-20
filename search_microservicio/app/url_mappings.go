package app

import (
	solrController "wesolr/controllers/solr"

	log "github.com/sirupsen/logrus"
)

func mapUrls() {

	router.GET("/search=:searchQuery", solrController.GetQuery)

	router.GET("/searchAll=:searchQuery", solrController.GetQueryAllFields)

	router.GET("/properties/:id", solrController.AddFromId)

	log.Info("Finishing mappings configurations")
}
