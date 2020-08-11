package v2

import (
	"github.com/gadzooks/weather-go-api/controller"
	v1Service "github.com/gadzooks/weather-go-api/service"
	v2Service "github.com/gadzooks/weather-go-api/service/v2"
	"github.com/gadzooks/weather-go-api/utils"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"net/http"
)

type PlaceController interface {
	FindLocations(w http.ResponseWriter, r *http.Request)
	SeedLocations(w http.ResponseWriter, r *http.Request)

	FindRegions(w http.ResponseWriter, r *http.Request)
	SeedRegions(w http.ResponseWriter, r *http.Request)
	CreateRegion(w http.ResponseWriter, r *http.Request)
	GetRegion(w http.ResponseWriter, r *http.Request)
	UpdateRegion(w http.ResponseWriter, r *http.Request)
	DeleteRegion(w http.ResponseWriter, r *http.Request)
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
	utils.SetLoggerWithRequestId(r.Context())
	log.Info().Msg("FindLocations")
	resp, err := ctrl.v2Svc.GetLocations()
	controller.HandleServiceResponse(w, resp, err)
}

// Get all regions
func (ctrl PlaceControllerImpl) FindRegions(w http.ResponseWriter, r *http.Request) {
	utils.SetLoggerWithRequestId(r.Context())
	log.Info().Msg("FindRegions")
	resp, err := ctrl.v2Svc.GetRegions()
	controller.HandleServiceResponse(w, resp, err)
}

// Create a new Region
func (ctrl PlaceControllerImpl) CreateRegion(w http.ResponseWriter, r *http.Request) {
	utils.SetLoggerWithRequestId(r.Context())
	log.Info().Msg("CreateRegion")
	vars := mux.Vars(r)
	regionData := map[string]string{
		"id":          vars["id"],
		"name":        vars["name"],
		"description": vars["description"],
	}
	resp, err := ctrl.v2Svc.CreateRegion(regionData)
	controller.HandleServiceResponse(w, resp, err)
}

// Get Region from db
func (ctrl PlaceControllerImpl) GetRegion(w http.ResponseWriter, r *http.Request) {
	utils.SetLoggerWithRequestId(r.Context())
	// to get logs with the request details :
	// GetRegion reqId=3497b9a8-747c-423a-8d8f-41559476795c request="GET /v2/region/foobar HTTP/1.1\r\nHost: localhost:8080\r\nAccept: */*\r\nAccept-Encoding: gzip, deflate, br\r\nConnection: keep-alive\r\nPostman-Token: 6db04fcc-9bd4-4503-9df2-0d0887401f9c\r\nUser-Agent: PostmanRuntime/7.26.2\r\n\r\n"
	// requestDump, _ := httputil.DumpRequest(r, false)
	// log.Logger = log.With().Str("request", string(requestDump)).Logger()
	log.Info().Msg("GetRegion")
	vars := mux.Vars(r)
	key := vars["id"]
	resp, err := ctrl.v2Svc.GetRegion(key)
	controller.HandleServiceResponse(w, resp, err)
}

// Update region
func (ctrl PlaceControllerImpl) UpdateRegion(w http.ResponseWriter, r *http.Request) {
	utils.SetLoggerWithRequestId(r.Context())
	log.Info().Msg("UpdateRegion")
	vars := mux.Vars(r)
	regionData := map[string]string{
		"id":          vars["id"],
		"name":        vars["name"],
		"description": vars["description"],
	}

	resp, err := ctrl.v2Svc.UpdateRegion(regionData)
	controller.HandleServiceResponse(w, resp, err)
}

// Delete a region
func (ctrl PlaceControllerImpl) DeleteRegion(w http.ResponseWriter, r *http.Request) {
	utils.SetLoggerWithRequestId(r.Context())
	log.Info().Msg("DeleteRegion")
	vars := mux.Vars(r)
	key := vars["id"]
	resp, err := ctrl.v2Svc.DeleteRegion(key)
	controller.HandleServiceResponse(w, resp, err)
}

// Seed regions with default data
func (ctrl PlaceControllerImpl) SeedRegions(w http.ResponseWriter, r *http.Request) {
	utils.SetLoggerWithRequestId(r.Context())
	log.Info().Msg("SeedRegions")
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
	utils.SetLoggerWithRequestId(r.Context())
	log.Info().Msg("SeedLocations")
	locations, err := ctrl.v1Svc.GetLocations()
	if err != nil {
		controller.HandleServiceResponse(w, locations, err)
	} else {
		resp, err := ctrl.v2Svc.SeedLocations(locations)
		controller.HandleServiceResponse(w, resp, err)
	}
}
