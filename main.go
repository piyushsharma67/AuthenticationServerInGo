package main

import (
	"chat-application/routes"
	"chat-application/server"
	"chat-application/store"
	"os"

	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
)


func main(){

	// *************http server implementation***************

	mongoUrl:=os.Getenv("MONGO_CONNECTION_STRING")
	httpPort:=os.Getenv("HTTP_PORT")
	
	store:=store.New(mongoUrl)

	fmt.Println("Creating server!!")
	server:=server.New(store)

	fmt.Println("Creating routes")
	routes:=routes.InitRoutes(server)

	fmt.Println("starting server on "+httpPort)
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization", "Content-Disposition"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "OPTIONS", "DELETE"})
	
	err:=http.ListenAndServe(fmt.Sprintf(":%s",httpPort),handlers.CORS(headersOk,originsOk,methodsOk)(routes))

	if(err!=nil){
		log.Fatal(err)
	}

	fmt.Println("running server!!")
}