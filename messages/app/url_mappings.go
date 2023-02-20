package app

import (
	messageController "messages/controllers/message"

	log "github.com/sirupsen/logrus"
)

func mapUrls() {
	// Messages Mapping
	router.GET("/messages/:id", messageController.GetMessageById)
	router.GET("/properties/:id/messages", messageController.GetMessagesByPropertyId)

	router.DELETE("/messages/:id", messageController.DeleteMessage)
	router.POST("/message", messageController.MessageInsert)

	log.Info("Finishing mappings configurations")
}
