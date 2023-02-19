package services

import (
	messageClient "messages/clients/message"
	"messages/dto"
	model "messages/model"
	e "messages/utils/errors"
)

type messageService struct{}

type MessageServiceInterface interface {
	GetMessageById(id string) (dto.MessageDto, e.ApiError)
}

var (
	MessageService MessageServiceInterface
)

func init() {
	MessageService = &messageService{}
}

func (s *messageService) GetMessageById(id string) (dto.MessageDto, e.ApiError) {

	var message model.Message = messageClient.GetMessageById(id)

	var messageDto dto.MessageDto

	if message.Id.Hex() == "000000000000000000000000" {
		return messageDto, e.NewBadRequestApiError("message not found")
	}

	messageDto.PropertyId = message.PropertyId
	messageDto.UserId = message.UserId
	messageDto.Body = message.Body
	messageDto.CreatedAt = message.CreatedAt
	messageDto.Id = message.Id.Hex()

	return messageDto, nil
}
