package app

import (
	solrController "search/controllers/solr"

	log "github.com/sirupsen/logrus"
)

func mapUrls() {

	router.GET("/properties/:id", solrController.AddFromId)

	router.DELETE("/properties/:id", solrController.Delete)

	log.Info("Finishing mappings configurations")
}
