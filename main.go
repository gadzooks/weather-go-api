package main

import (
	"net/http"

	"github.com/gadzooks/weather-go-api/config"
)

func main() {
	r := config.NewRouter()

	config.AddAPISubRouterForPlaces(r)

	http.ListenAndServe(":80", r)
}
