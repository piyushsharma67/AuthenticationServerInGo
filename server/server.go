package server

import "chat-application/store"

type Server struct{
	store *store.Store 
}
func New(store *store.Store)*Server{

	server := &Server{}
	server.store=store

	return server
}