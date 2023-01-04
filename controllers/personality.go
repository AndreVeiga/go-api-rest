package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/AndreVeiga/go-api-rest/models"
	"github.com/gorilla/mux"
)

func GetAllPersonalities(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.GetAll())
}

func GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	personalities := models.GetAll()

	for _, personality := range personalities {
		if strconv.Itoa(personality.Id) == id {

		}
	}
}
