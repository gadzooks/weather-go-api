package config

import (
	"github.com/gadzooks/weather-go-api/client"
	v2Client "github.com/gadzooks/weather-go-api/client/v2"
	"github.com/gadzooks/weather-go-api/controller"
	v2Controller "github.com/gadzooks/weather-go-api/controller/v2"
	"github.com/gadzooks/weather-go-api/service"
	v2Service "github.com/gadzooks/weather-go-api/service/v2"
	"github.com/gorilla/mux"
)

// NewRouter creates new base router
func NewRouter() *mux.Router {
	r := mux.NewRouter()

	// todo handle static assets
	return r
}

const storageDataDir = "data"

func AddV2APISubRouterForPlaces(base *mux.Router) {
	// restrict to all urls under /api
	api := base.PathPrefix("/v2").Subrouter()

	// create repo object
	v1Client := client.NewStorageClient(storageDataDir)
	v2Client := v2Client.NewStorageClient("")
	// create service object
	v1Service := service.NewPlaceService(v1Client)
	v2Service := v2Service.NewPlaceService(v2Client)
	// create controller object
	placesCtrl := v2Controller.NewPlaceController(v1Service, v2Service)

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
	api.HandleFunc("/locations", placesCtrl.FindLocations).Methods("GET")

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

	// SeedRegions swagger:route POST /regions regions findRegions
	// Seeds a region set with default values
	//
	// Consumes:
	// - application/json
	//
	// Produces:
	// - application/json
	//
	// Responses:
	// 200: []region
	api.HandleFunc("/regions", placesCtrl.SeedRegions).Methods("POST")
	/*
		a.Router.HandleFunc("/products", a.getProducts).Methods("GET")
		a.Router.HandleFunc("/product", a.createProduct).Methods("POST")
		a.Router.HandleFunc("/product/{id:[0-9]+}", a.getProduct).Methods("GET")
		a.Router.HandleFunc("/product/{id:[0-9]+}", a.updateProduct).Methods("PUT")
		a.Router.HandleFunc("/product/{id:[0-9]+}", a.deleteProduct).Methods("DELETE")

	*/
}

func AddAPISubRouterForPlaces(base *mux.Router) {
	// restrict to all urls under /api
	api := base.PathPrefix("/v1").Subrouter()

	// create repo object
	client := client.NewStorageClient(storageDataDir)
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
	api.HandleFunc("/locations", placesCtrl.FindLocations).Methods("GET")

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
