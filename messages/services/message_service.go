package services

import (
	messageCliente "messages/clients/message"
	"messages/dto"
	"messages/model"
	e "messages/utils/errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type messageService struct{}

type messageServiceInterface interface {
	GetMessageByReceiver(token string) (dto.MessagesDto, e.ApiError)
	GetMessages() (dto.MessagesDto, e.ApiError)
	InsertMessage(messageDto dto.MessageDto) (dto.MessageDto, e.ApiError)
}

var (
	MessageService messageServiceInterface
)

func init() {
	MessageService = &messageService{}
}

func (s *messageService) GetMessages() (dto.MessagesDto, e.ApiError) {

	var messages model.Messages = messageCliente.GetMessages()
	var messagesDto dto.MessagesDto

	for _, message := range messages {
		var messageDto dto.MessageDto
		messageDto.Id = message.Id
		messageDto.Receiver = message.Receiver
		messageDto.Sender = message.Sender
		messageDto.Message = message.Message
		messageDto.Date = message.Date

		messagesDto = append(messagesDto, messageDto)
	}

	return messagesDto, nil
}

func (s *messageService) InsertMessage(messageDto dto.MessageDto) (dto.MessageDto, e.ApiError) {

	var message model.Message

	message.Message = messageDto.Message
	message.Receiver = messageDto.Receiver
	message.Sender = messageDto.Sender
	message.Date = time.Now()
	message = messageCliente.InsertMessage(message)

	messageDto.Id = message.Id

	return messageDto, nil
}

func (s *messageService) GetMessageByReceiver(token string) (dto.MessagesDto, e.ApiError) {

	var IdReceiver float64
	tkn, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) { return jwtKey, nil })

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, e.NewUnauthorizedApiError("Unauthorized")
		}
		return nil, e.NewUnauthorizedApiError("Unauthorized")
	}

	if !tkn.Valid {
		return nil, e.NewUnauthorizedApiError("Unauthorized")

	}
	if claims, ok := tkn.Claims.(jwt.MapClaims); ok && tkn.Valid {

		IdReceiver = (claims["id_receiver"].(float64))

	} else {
		return nil, e.NewUnauthorizedApiError("Unauthorized")
	}
	var IdReceiverX int = int(IdReceiver)
	var messages model.Messages = messageCliente.GetMessageByReceiver(IdReceiverX)
	var messagesDto dto.MessagesDto

	if len(messages) == 0 {
		return messagesDto, e.NewBadRequestApiError("message not found")
	}

	for _, message := range messages {
		var messageDto dto.MessageDto
		messageDto.Id = message.Id
		messageDto.Receiver = message.Receiver
		messageDto.Sender = message.Sender
		messageDto.Message = message.Message
		messageDto.Date = message.Date

		messagesDto = append(messagesDto, messageDto)
	}

	return messagesDto, nil
}
