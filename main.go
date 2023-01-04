package main

import (
	"github.com/AndreVeiga/go-api-rest/database"
	"github.com/AndreVeiga/go-api-rest/routes"
)

func main() {
	database.ConnectDatabase()

	routes.HandleRequets()
}
