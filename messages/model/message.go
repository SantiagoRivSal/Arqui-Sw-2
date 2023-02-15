package model

import (
	"time"
)

type Message struct {
	Id       int       `gorm:"primaryKey"`
	Receiver int       `gorm:"type:int(150);not null"`
	Sender   int       `gorm:"type:int(150);not null"`
	Message  string    `gorm:"type:varchar(500);not null"`
	Date     time.Time `gorm:"not null"`
}

type Messages []Message
