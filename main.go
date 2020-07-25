package main

import (
	"net/http"

	"github.com/gadzooks/weather-api-go/config"
)

func main() {
	r := config.NewRouter()

	config.AddAPISubRouterForPlaces(r)

	http.ListenAndServe(":80", r)
}
