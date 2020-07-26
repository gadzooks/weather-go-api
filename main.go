package main

import (
	"github.com/gadzooks/weather-go-api/middleware"
	"log"
	"net/http"

	"github.com/gadzooks/weather-go-api/config"
)

func main() {
	r := config.NewRouter()

	config.AddAPISubRouterForPlaces(r)

	// add middleware
	handler := middleware.SetupGlobalMiddleware(r)
	handler = middleware.NewLogger(handler)

	log.Println("starting server at 8080")
	err := http.ListenAndServe(":8080", handler)
	if err != nil {
		log.Fatalf("error running server : %v", err)
	}

}
