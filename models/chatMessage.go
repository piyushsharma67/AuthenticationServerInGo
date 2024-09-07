package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ChatMessage struct{
	ID primitive.ObjectID `json:"_id" bson:"_id"`
	ChatRoomID primitive.ObjectID `json:"chat_room_id" bson:"chat_room_id"`
	SenderID primitive.ObjectID `json:"sender_id" bson:"sender_id"`
	Message string `json:"message" bson:"message"`
	CreatedAt primitive.DateTime `json:"created_at" bson:"created_at"`
	RecieverID primitive.ObjectID `json:"reciever_id" bson:"reciever_id"`
} 