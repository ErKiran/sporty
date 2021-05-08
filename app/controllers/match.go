package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"play-it/app/models"
	"play-it/app/responses"
)

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

	js, _ := json.MarshalIndent(data, "", " ")
	fmt.Println("js", string(js))
}
