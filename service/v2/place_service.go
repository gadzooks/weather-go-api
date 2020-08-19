package v2

import (
	v2Client "github.com/gadzooks/weather-go-api/client/v2"
	"github.com/gadzooks/weather-go-api/model"
	"github.com/rs/zerolog/log"
)

// PlaceService is responsible for querying locations and regions
type PlaceService interface {
	GetLocations() ([]model.Location, error)
	SeedLocations([]model.Location) ([]model.Location, error)

	SeedRegions([]model.Region) ([]model.Region, error)
	GetRegions() ([]model.Region, error)

	// CRUD operations for region model
	GetRegion(string) (model.Region, error)
	CreateRegion(region model.Region) (model.Region, error)
	UpdateRegion(map[string]string) error
	DeleteRegion(string) error
}

// PlaceServiceImpl implements PlaceService
type PlaceServiceImpl struct {
	client v2Client.StorageClient
}

func (r PlaceServiceImpl) GetLocations() ([]model.Location, error) {
	var results []model.Location
	loc, err := r.client.QueryLocations()
	if err != nil {
		return results, err
	}
	for _, m := range loc {
		results = append(results, m)
	}
	return results, nil
}

func (r PlaceServiceImpl) SeedLocations(data []model.Location) ([]model.Location, error) {
	var inserted []model.Location
	log.Info().Msgf("SeedLocations with : %v", data)
	err := r.client.DeleteAllLocations()
	if err != nil {
		return inserted, err
	}
	for _, location := range data {
		log.Info().Msgf("insert location : %v\n", location)
		resp, err := r.client.CreateLocation(location)
		if err != nil {
			log.Info().Msgf("error inserting location : %v\n", err)
		} else {
			inserted = append(inserted, location)
			log.Info().Msgf("inserted location with id : %v", resp.ID.Hex())
		}
	}
	return inserted, nil
}

var InvalidInputError = v2Client.InvalidInputError
var NotFoundError = v2Client.NotFoundError

func (r PlaceServiceImpl) GetRegion(id string) (model.Region, error) {
	return r.client.FindRegion(id)
}

func (r PlaceServiceImpl) UpdateRegion(region map[string]string) error {
	params := make(map[string]string)
	params["id"] = region["id"]
	// FIXME : support name update, which needs updating location docs too
	params["description"] = region["description"]

	return r.client.UpdateRegion(params)
}

func (r PlaceServiceImpl) CreateRegion(region model.Region) (model.Region, error) {
	return r.client.CreateRegion(region)
}

func (r PlaceServiceImpl) DeleteRegion(id string) error {
	return r.client.DeleteRegion(id)
}

func (r PlaceServiceImpl) GetRegions() ([]model.Region, error) {
	var results []model.Region
	loc, err := r.client.QueryRegions()
	if err != nil {
		return results, err
	}
	for _, m := range loc {
		results = append(results, m)
	}
	return results, nil
}

func (r PlaceServiceImpl) SeedRegions(data []model.Region) ([]model.Region, error) {
	var inserted []model.Region
	log.Info().Msgf("SeedRegions with : %v", data)
	err := r.client.DeleteAllRegions()
	if err != nil {
		return inserted, err
	}
	for _, region := range data {
		log.Info().Msgf("insert region : %v\n", region)
		resp, err := r.client.CreateRegion(region)
		if err != nil {
			log.Info().Msgf("error inserting region : %v\n", err)
		} else {
			inserted = append(inserted, region)
			log.Info().Msgf("inserted region with id : %v", resp.ID.Hex())
		}
	}
	return inserted, nil
}

// NewPlaceService creates new instance of PlaceService
func NewPlaceService(client v2Client.StorageClient) PlaceService {
	return &PlaceServiceImpl{client: client}
}
