package main

import (
	"log"
	"net/http"

	"github.com/gadzooks/weather-go-api/config"
)

func main() {
	r := config.NewRouter()

	config.AddAPISubRouterForPlaces(r)

	log.Println("starting server at 8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalf("error running server : %v", err)
	}

}
