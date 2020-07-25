package config

import (
	"github.com/gadzooks/weather-go-api/client"
	"github.com/gadzooks/weather-go-api/controller"
	"github.com/gadzooks/weather-go-api/service"
	"github.com/gorilla/mux"
)

// NewRouter creates new base router
func NewRouter() *mux.Router {
	r := mux.NewRouter()

	// todo handle static assets
	return r
}

func AddAPISubRouterForPlaces(base *mux.Router) {
	// restrict to all urls under /api
	api := base.PathPrefix("/api").Subrouter()

	// create repo object
	client := client.NewStorageClient("data")
	// create service object
	svc := service.NewPlaceService(client)
	// create controller object
	placesCtrl := controller.NewPlaceController(svc)

	// FindLocations swagger:route GET /locations locations findLocations
	//
	// Finds a location set
	//
	// Consumes:
	// - application/json
	//
	// Produces:
	// - application/json
	//
	// Responses:
	// 200: []location
	api.HandleFunc("/locations", placesCtrl.FindLocations).Methods("GET") // FindLocations swagger:route GET /locations locations findSamples

	// FindRegions swagger:route GET /regions regions findRegions
	// Finds a region set
	//
	// Consumes:
	// - application/json
	//
	// Produces:
	// - application/json
	//
	// Responses:
	// 200: []region
	api.HandleFunc("/regions", placesCtrl.FindRegions).Methods("GET")

	/*
		a.Router.HandleFunc("/products", a.getProducts).Methods("GET")
		a.Router.HandleFunc("/product", a.createProduct).Methods("POST")
		a.Router.HandleFunc("/product/{id:[0-9]+}", a.getProduct).Methods("GET")
		a.Router.HandleFunc("/product/{id:[0-9]+}", a.updateProduct).Methods("PUT")
		a.Router.HandleFunc("/product/{id:[0-9]+}", a.deleteProduct).Methods("DELETE")

	*/
}
