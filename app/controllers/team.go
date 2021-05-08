package controllers

import (
	"net/http"
	"play-it/app/models"
	"play-it/app/responses"
)

func (server *Server) GetTeams(w http.ResponseWriter, r *http.Request) {
	teams := models.Team{}

	res, err := teams.FindAll(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	responses.JSON(w, 200, res)
}
