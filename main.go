package main

import (
	"log"
	"net/http"

	"github.com/johnnyFR26/go.api/routes"
)

func main() {
	router := routes.SetupRouter()

	log.Println("Server running on http://localhost:8000")
	http.ListenAndServe(":8000", router)
}
