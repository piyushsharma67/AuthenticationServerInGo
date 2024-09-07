package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ChatRoom struct{
	ID primitive.ObjectID  `json:"_id" bson:"_id"`
	Participants []string `json:"participants" bson:"participants"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	
}