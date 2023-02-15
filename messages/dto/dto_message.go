package dto

import (
	"time"
)

type MessageDto struct {
	Receiver int       `json:"receiver"`
	Sender   int       `json:"sender"`
	Message  string    `json:"message"`
	Date     time.Time `json:"date"`
	Id       int       `json:"id"`
}

type MessagesDto []MessageDto
