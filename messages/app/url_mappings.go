package app

import (
	messageController "messages/controllers/message"

	log "github.com/sirupsen/logrus"
)

func mapUrls() {

	// messages Mapping
	router.GET("/message/:receiver", messageController.GetMessageByReceiver)
	router.GET("/message", messageController.GetMessages)
	router.POST("/message", messageController.InserMessage)

	log.Info("Finishing mappings configurations")
}
