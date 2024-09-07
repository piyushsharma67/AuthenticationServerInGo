package store

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Store struct{
	DB *mongo.Database 
}


func New(mongoUrl string) *Store{

	store:= &Store{}

	clientOption:=options.Client().ApplyURI(mongoUrl)
	client, err := mongo.Connect(context.Background(), clientOption)

	if(err!=nil){
		log.Fatal(err)
	}

	store.DB=client.Database("chat-application")

	return store

}