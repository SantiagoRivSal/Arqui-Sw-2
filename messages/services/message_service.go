package services

import (
	"messages/dto"
	e "messages/utils/errors"
)

type massageService struct{}

type messageServiceInterface interface {
	GetMessageByReceiver(id int) (dto.MessageDto, e.ApiError)
	GetMessages() (dto.MessagesDto, e.ApiError)
	InserMessage(messageDto dto.MessageDto) (dto.MessageDto, e.ApiError)
}

var (
	MessageService messageServiceInterface
)
