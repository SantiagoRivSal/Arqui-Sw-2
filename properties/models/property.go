package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Address struct {
}

type Property struct {
	Id          primitive.ObjectID `bson:"_id"`
	Tittle      string             `bson:"tittle"`
	Description string             `bson:"description"`
	Size        int                `bson:"size"`
	Rooms       int                `bson:"rooms"`
	Bathrooms   int                `bson:"bathrooms"`
	Service     string             `bson:"service"`
	Price       int                `bson:"price"`
	Image       string             `bson:"image"`
	UserId      int                `bson:"userid"`
	Street      string             `bson:"street"`
	City        string             `bson:"city"`
	State       string             `bson:"state"`
	Country     string             `bson:"country"`
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
