package main

import (
	"log"
	"net/http"

	"github.com/gadzooks/weather-go-api/config"
)

func main() {
	r := config.NewRouter()

	config.AddAPISubRouterForPlaces(r)

	err := http.ListenAndServe(":80", r)
	if err != nil {
		log.Fatalf("error running server : %v", err)
	}

}
