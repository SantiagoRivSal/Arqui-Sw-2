package dto

type MessageDto struct {
	Receiver string `json:"receiver"`
	Sender   string `json:"sender"`
	Message  string `json:"message"`
	Date     string `json:"date"`
	Id       int    `json:"id"`
}

type MessagesDto []MessageDto
