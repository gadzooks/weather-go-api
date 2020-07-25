package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gadzooks/weather-api-go/service"
)

type PlaceController interface {
	FindLocations(w http.ResponseWriter, r *http.Request)
	FindRegions(w http.ResponseWriter, r *http.Request)
}

type PlaceControllerImpl struct {
	svc service.PlaceService
}

func NewPlaceController(svc service.PlaceService) PlaceController {
	return &PlaceControllerImpl{svc: svc}
}

// Get all locations
func (ctrl PlaceControllerImpl) FindLocations(w http.ResponseWriter, r *http.Request) {
	resp, err := ctrl.svc.GetLocations()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	js, err := json.MarshalIndent(resp, "", "")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(js)
}

// Get all regions
func (ctrl PlaceControllerImpl) FindRegions(w http.ResponseWriter, r *http.Request) {
	resp, err := ctrl.svc.GetRegions()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	js, err := json.MarshalIndent(resp, "", "")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(js)

}
