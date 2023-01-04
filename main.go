package main

import (
	"fmt"

	"github.com/AndreVeiga/go-api-rest/routes"
)

func main() {
	fmt.Println("Starting server in Golang")
	routes.HandleRequets()
}
