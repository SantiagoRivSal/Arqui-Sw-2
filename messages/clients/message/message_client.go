package message

import (
	"messages/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetMessages() model.Messages {
	var messages model.Messages
	Db.Find(&messages)

	log.Debug("Messages: ", messages)

	return messages
}

func InsertMessage(message model.Message) model.Message {
	result := Db.Create(&message)

	if result.Error != nil {
		//TODO Manage Errors
		log.Error("")
	}
	log.Debug("Message Created: ", message.Id)
	return message
}
func GetUserByReceiver(receiver int) model.Messages {
	var messages model.Messages

	Db.Where("receiver = ?", receiver).First(&messages)
	log.Debug("Receiver: ", receiver)
	log.Debug("Message: ", messages)
	return messages
}
