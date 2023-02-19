package dto

type MessageDto struct {
	Id         int    `json:"id"`
	UserId     int    `json:"userid"`
	PropertyId string `json:"propertyid"`
	Body       string `json:"body"`
	CreatedAt  string `json:"created_at"`
}

type MessagesDto []MessageDto
