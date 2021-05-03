package api

import (
	"play-it/models"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	DB     *mongo.Client
	Router *mux.Router
}

func New() *Server {
	s := &Server{}
	s.DB = models.SetUp()
	s.Router = mux.NewRouter()
	return s
}
