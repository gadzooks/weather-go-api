package v2

import (
	"github.com/gadzooks/weather-go-api/controller"
	v1Service "github.com/gadzooks/weather-go-api/service"
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
	v1Svc v1Service.PlaceService
	v2Svc v2Service.PlaceService
}

func NewPlaceController(v1 v1Service.PlaceService, v2 v2Service.PlaceService) PlaceController {
	return &PlaceControllerImpl{v2Svc: v2, v1Svc: v1}
}

// Get all locations
func (ctrl PlaceControllerImpl) FindLocations(w http.ResponseWriter, r *http.Request) {
	// FIXME update to v2
	resp, err := ctrl.v1Svc.GetLocations()
	controller.HandleServiceResponse(w, resp, err)
}

// Get all regions
func (ctrl PlaceControllerImpl) FindRegions(w http.ResponseWriter, r *http.Request) {
	// FIXME update to v2
	resp, err := ctrl.v1Svc.GetRegions()
	controller.HandleServiceResponse(w, resp, err)
}

// Seed regions with default data
func (ctrl PlaceControllerImpl) SeedRegions(w http.ResponseWriter, r *http.Request) {
	regions, err := ctrl.v1Svc.GetRegions()
	if err != nil {
		controller.HandleServiceResponse(w, regions, err)
	} else {
		resp, err := ctrl.v2Svc.SeedRegions(regions)
		controller.HandleServiceResponse(w, resp, err)
	}
}

// Seed locations with default data
func (ctrl PlaceControllerImpl) SeedLocations(w http.ResponseWriter, r *http.Request) {
	locations, err := ctrl.v1Svc.GetLocations()
	if err != nil {
		controller.HandleServiceResponse(w, locations, err)
	} else {
		resp, err := ctrl.v2Svc.SeedLocations(locations)
		controller.HandleServiceResponse(w, resp, err)
	}
}
