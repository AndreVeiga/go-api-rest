package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AndreVeiga/go-api-rest/controllers"
	"github.com/AndreVeiga/go-api-rest/middlewares"
	"github.com/AndreVeiga/go-api-rest/routes/enums"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequets() {
	r := mux.NewRouter()
	r.Use(middlewares.ContentTypeMiddleware)

	r.HandleFunc("/", controllers.Home).Methods(enums.GET)
	r.HandleFunc("/personalities", controllers.GetAllPersonalities).Methods(enums.GET)
	r.HandleFunc("/personalities/{id}", controllers.GetPersonalitiesById).Methods(enums.GET)
	r.HandleFunc("/personalities", controllers.CreatePersonalities).Methods(enums.POST)
	r.HandleFunc("/personalities/{id}", controllers.DeletePersonalitiesById).Methods(enums.DELETE)
	r.HandleFunc("/personalities/{id}", controllers.EditPersonalitiesById).Methods(enums.PUT)

	fmt.Println("Starting server in Golang")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
