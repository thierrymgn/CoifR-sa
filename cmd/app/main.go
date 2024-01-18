package main

import (
	"log"
	"net/http"

	"CoifResa/pkg/api"
)

func main() {
	r := api.NewRouter()

	log.Println("Le serveur écoute sur le port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
