package controllers

import (
	"log"
	"net/http"
	"play-it/app/middleware"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	DB     *mongo.Client
	Router *mux.Router
}

func (server *Server) setJSON(path string, next func(http.ResponseWriter, *http.Request), method string) {
	server.Router.HandleFunc(path, middleware.SetMiddlewareJSON(next)).Methods(method, "OPTIONS")
}

func (server *Server) InitRoutes() {
	server.setJSON("/players", server.GetPlayers, "GET")
}

func (server *Server) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
