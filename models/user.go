package models

import (
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SignUpRequestBody struct{
	Name string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignInRequestBody struct{
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserModel struct {
	ID primitive.ObjectID   `json:"_id" bson:"_id"`
	Name string `json:"name" bson:"name"`
	Email string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}

func(user *UserModel) Normalise() *UserModel{
	if(user.Email!=""){
		user.Email=strings.ToLower(user.Email)
	}

	return user
}

func(user *UserModel) PreInsert() *UserModel{
	var timeNow=time.Now().UTC()
	user.UpdatedAt=timeNow
	user.CreatedAt=timeNow

	return user
}