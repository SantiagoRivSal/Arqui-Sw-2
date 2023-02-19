package messageController

import (
	service "messages/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMessageById(c *gin.Context) {

	id := c.Param("parameters")

	messageDto, er := service.MessageService.GetMessageById(id)

	// Error del Insert
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusOK, messageDto)
}
