package store

import (
	"chat-application/models"
	"context"
	"fmt"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	User_Collection="UserCollection"
)

func (s *Store) AddUser(srb *models.SignUpRequestBody) (*models.UserModel,error){
	ctx := context.TODO()

	user:=&models.UserModel{
		ID: primitive.NewObjectID(),
		Email: srb.Email,
		Password: srb.Password,
		Name: srb.Name,
	}

	user.Normalise()
	user.PreInsert()

	_,error:=s.DB.Collection(User_Collection).InsertOne(ctx,user)

	if(error!=nil){
		return nil,error
	}
	
	return user,nil
}
















func (s *Store) GetUserByEmail(email string) (*models.UserModel,error){
	ctx := context.TODO()
	user:=&models.UserModel{}
	
	// Use bson.M to construct the filter
	filter := bson.M{"email": strings.ToLower(email)}

	err:=s.DB.Collection(User_Collection).FindOne(ctx,filter).Decode(user)
	fmt.Println(err,user)
	if err != nil {
        if err == mongo.ErrNoDocuments {
            // No user found with the given email
            return nil, nil
        }
        // An error occurred while querying the database
        return nil, err
    }

    // User found
    return user, nil
}