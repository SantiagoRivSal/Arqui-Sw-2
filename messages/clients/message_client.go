package message

import (
	"messages/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetMessages() model.messages {
	var messages model.messages
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
func GetUserByReceiver(receiver string) model.Message {
	var message model.Message

	Db.Where("receiver = ?", receiver).First(&message)
	log.Debug("Receiver: ", message.receiver)

	return message
}
