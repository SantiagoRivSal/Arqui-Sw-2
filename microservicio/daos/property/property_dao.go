package property

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	model "microservicio/models"
	"microservicio/utils/db"
)

func GetById(id string) model.Property {
	var property model.Property
	db := db.MongoDb
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println(err)
		return property
	}
	err = db.Collection("properties").FindOne(context.TODO(), bson.D{{"_id", objID}}).Decode(&property)
	if err != nil {
		fmt.Println(err)
		return property
	}
	return property

}

func Insert(property model.Property) model.Property {
	db := db.MongoDb
	insertProperty := property
	insertProperty.Id = primitive.NewObjectID()
	_, err := db.Collection("properties").InsertOne(context.TODO(), &insertProperty)

	if err != nil {
		fmt.Println(err)
		return property
	}
	fmt.Println("id insertada: ", insertProperty.Id)
	property.Id = insertProperty.Id
	return property
}
