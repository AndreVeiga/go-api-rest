package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AndreVeiga/go-api-rest/controllers"
	"github.com/gorilla/mux"
)

func HandleRequets() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home).Methods("GET")
	r.HandleFunc("/personalities", controllers.GetAllPersonalities).Methods("GET")
	r.HandleFunc("/personalities/{id}", controllers.GetById).Methods("GET")

	fmt.Println("Starting server in Golang")
	log.Fatal(http.ListenAndServe(":8000", r))
}
