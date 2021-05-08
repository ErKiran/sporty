package controllers

import (
	"net/http"
	"play-it/app/models"
	"play-it/app/responses"

	"github.com/gorilla/mux"
)

func (server *Server) GetPlayers(w http.ResponseWriter, r *http.Request) {
	players := models.Player{}

	res, err := players.FindAll(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	responses.JSON(w, 200, res)
}

func (server *Server) GetPlayersByTeamId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	players := models.Player{}

	res, err := players.FindPlayersOfTeam(server.DB, vars["id"])
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	responses.JSON(w, 200, res)

}
