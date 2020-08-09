package controller

import (
	"encoding/json"
	"github.com/gadzooks/weather-go-api/service"
	"net/http"
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
	HandleServiceResponse(w, resp, err)
}

// Get all regions
func (ctrl PlaceControllerImpl) FindRegions(w http.ResponseWriter, r *http.Request) {
	resp, err := ctrl.svc.GetRegions()
	HandleServiceResponse(w, resp, err)
}

func HandleServiceResponse(w http.ResponseWriter, resp interface{}, err error) {
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
