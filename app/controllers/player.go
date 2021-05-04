package controllers

import (
	"net/http"
	"play-it/app/models"
	"play-it/app/responses"
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
