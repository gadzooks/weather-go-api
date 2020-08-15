package controller

import (
	"encoding/json"
	"errors"
	"github.com/gadzooks/weather-go-api/service"
	v2 "github.com/gadzooks/weather-go-api/service/v2"
	"github.com/gadzooks/weather-go-api/utils"
	"github.com/rs/zerolog/log"
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
	utils.SetLoggerWithRequestId(r.Context())
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
		log.Error().Msg(err.Error())
		js, _ := json.MarshalIndent(err.Error(), "", "")
		if errors.Is(err, v2.InvalidInputError) {
			log.Info().Msg(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			if js != nil {
				_, _ = w.Write(js)
			}
			return
		} else if errors.Is(err, v2.NotFoundError) {
			log.Info().Msg(err.Error())
			w.WriteHeader(http.StatusNotFound)
			if js != nil {
				_, _ = w.Write(js)
			}
			return
		}
		log.Error().Msg(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	if resp != nil {
		js, err := json.MarshalIndent(resp, "", "")
		if err != nil {
			log.Error().Msg(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if js != nil {
			_, _ = w.Write(js)
		}
	}

}
