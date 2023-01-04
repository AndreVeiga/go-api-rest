package controllers

import (
	"encoding/json"
	"log"
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

func GetPersonalitiesById(w http.ResponseWriter, r *http.Request) {
	var personality models.Personality
	database.DB.First(&personality, mux.Vars(r)["id"])

	json.NewEncoder(w).Encode(personality)
}

func CreatePersonalities(w http.ResponseWriter, r *http.Request) {
	var personality models.Personality
	if err := json.NewDecoder(r.Body).Decode(&personality); err != nil {
		log.Panic(err.Error())
		return
	}

	database.DB.Create(&personality)
	json.NewEncoder(w).Encode(&personality)
}

func DeletePersonalitiesById(w http.ResponseWriter, r *http.Request) {
	var personality models.Personality
	database.DB.Delete(&personality, mux.Vars(r)["id"])
}

func EditPersonalitiesById(w http.ResponseWriter, r *http.Request) {
	var personality models.Personality
	database.DB.First(&personality, mux.Vars(r)["id"])

	json.NewDecoder(r.Body).Decode(&personality)
	database.DB.Save(&personality)

	json.NewEncoder(w).Encode(&personality)
}
