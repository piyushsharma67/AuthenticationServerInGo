package routes

import (
	"chat-application/server"

	"github.com/gorilla/mux"
)

func InitRoutes(server *server.Server) *mux.Router{
	r:=mux.NewRouter()
	r.HandleFunc("/signup",server.Signup).Methods("POST")
	r.HandleFunc("/signin",server.Signin)

	return r
}