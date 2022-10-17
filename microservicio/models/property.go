package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Address struct {
	Street  string
	City    string
	State   string
	Country string
}

type Property struct {
	Id          primitive.ObjectID `bson:"_id"`
	Tittle      string             `bson:"text"`
	Description string             `bson:"text"`
	Address     Address            `bson:"inline"`
	Size        int
	Rooms       int
	Bathrooms   int
	Service     string `bson:"text"`
	Price       int
}

/*
{
	"tittle": "",
	"description": "",
	"country": "",
	"city": "",
	"size":  ,
    "rooms": ,
    "bathrooms": ,
    "service": "", ---> Puede ir como Alquiler o Venta
    "price":,
    "image": ""
}
*/
