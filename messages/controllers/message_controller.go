package userController

import (
	"messages/dto"
	service "messages/services"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetMessageByReceiver(c *gin.Context) {
	var messagesDto dto.MessageDto
	token := c.Param("receiver")
	messagesDto, err := service.MessageService.GetMessageByReceiver(token)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, messagesDto)
}

func GetMessages(c *gin.Context) {
	var messagesDto dto.MessagesDto
	messagesDto, err := service.MessageService.GetMessages()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, messagesDto)
}

func InserMessage(c *gin.Context) {
	var messageDto dto.MessageDto
	err := c.BindJSON(&messageDto)

	// Error Parsing json param
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	messageDto, er := service.MessageService.InserMessage(messageDto)
	// Error del Insert
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, messageDto)
}
