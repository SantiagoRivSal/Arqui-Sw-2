package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Address struct {
	Street  string `bson:"street"`
	City    string `bson:"city"`
	State   string `bson:"state"`
	Country string `bson:"country"`
}

type Property struct {
	Id          primitive.ObjectID `bson:"_id"`
	Tittle      string             `bson:"tittle"`
	Description string             `bson:"description"`
	Address     Address            `bson:"inline"`
	Size        int                `bson:"size"`
	Rooms       int                `bson:"rooms"`
	Bathrooms   int                `bson:"bathrooms"`
	Service     string             `bson:"service"`
	Price       int                `bson:"price"`
	Image       string             `bson:"image"`
}
type Properties []Property

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
