package app

import (
	"log"
	"play-it/app/controllers"
	"play-it/app/models"

	"github.com/gorilla/mux"
)

func New() *controllers.Server {
	s := &controllers.Server{}
	db, err := models.SetUp()
	if err != nil {
		log.Fatal(err)
	}
	s.DB = db
	s.Router = mux.NewRouter()
	return s
}

func Run() {
	s := New()
	s.InitRoutes()
	s.Run(":8080")
}
