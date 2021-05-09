package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"play-it/app/models"
	"play-it/app/responses"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MatchResponse struct {
	Id           string    `json:"id"`
	Title        string    `json:"title"`
	Location     string    `json:"location"`
	StartTime    time.Time `json:"startTime"`
	MatchFee     int       `json:"matchFee"`
	PlayerStatus string    `json:"playerStatus"`
}

func (server *Server) CreateMatch(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	data := models.Match{}

	err = json.Unmarshal(body, &data)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = data.Create(server.DB)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	responses.JSON(w, 200, nil)
}

func (server *Server) GetMatchOfPlayer(w http.ResponseWriter, r *http.Request) {

	data := models.Match{}

	id := r.Header.Get("id")
	oid, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	matches, err := data.Find(server.DB, oid.Hex())

	var response []MatchResponse

	for _, match := range matches {
		var matchFee int
		var status string
		for _, player := range match.Players {
			if player.PlayerID == oid {
				matchFee = player.MatchFee
				status = player.Status
			}

			if status == "" {
				status = "Undecided"
			}
		}
		response = append(response, MatchResponse{
			Id:           match.ID.Hex(),
			Title:        match.Title,
			Location:     match.Location,
			StartTime:    match.StartDate,
			PlayerStatus: status,
			MatchFee:     matchFee,
		})
	}

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	responses.JSON(w, 200, response)
}
