package message

import (
	"context"
	"fmt"
	model "messages/model"
	"messages/utils/db"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetMessageById(id string) model.Message {
	var message model.Message
	db := db.MongoDb
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println(err)
		return message
	}
	err = db.Collection("messages").FindOne(context.TODO(), bson.D{{"_id", objID}}).Decode(&message)
	if err != nil {
		fmt.Println(err)
		return message
	}
	return message

}
