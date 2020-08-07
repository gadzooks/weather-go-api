package v2

import (
	v2Service "github.com/gadzooks/weather-go-api/service/v2"
	"net/http"
)

type PlaceController interface {
	FindLocations(w http.ResponseWriter, r *http.Request)
	FindRegions(w http.ResponseWriter, r *http.Request)
	SeedLocations(w http.ResponseWriter, r *http.Request)
	SeedRegions(w http.ResponseWriter, r *http.Request)
}

type PlaceControllerImpl struct {
	svc v2Service.PlaceService
}

func NewPlaceController(svc v2Service.PlaceService) PlaceController {
	return &PlaceControllerImpl{svc: svc}
}

// Get all locations
func (ctrl PlaceControllerImpl) FindLocations(w http.ResponseWriter, r *http.Request) {

}

// Get all regions
func (ctrl PlaceControllerImpl) FindRegions(w http.ResponseWriter, r *http.Request) {

}

// Seed locations with default data
func (ctrl PlaceControllerImpl) SeedLocations(w http.ResponseWriter, r *http.Request) {

}

// Seed regions with default data
func (ctrl PlaceControllerImpl) SeedRegions(w http.ResponseWriter, r *http.Request) {

}
