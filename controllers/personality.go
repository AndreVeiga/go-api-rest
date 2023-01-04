package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/AndreVeiga/go-api-rest/database"
	"github.com/AndreVeiga/go-api-rest/models"
	"github.com/gorilla/mux"
)

func GetAllPersonalities(w http.ResponseWriter, r *http.Request) {
	var personalities []models.Personality
	database.DB.Find(&personalities)

	json.NewEncoder(w).Encode(personalities)
}

func GetById(w http.ResponseWriter, r *http.Request) {
	var personality models.Personality
	database.DB.First(&personality, mux.Vars(r)["id"])

	json.NewEncoder(w).Encode(personality)
}
