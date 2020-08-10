package v2

import (
	v2Client "github.com/gadzooks/weather-go-api/client/v2"
	"github.com/gadzooks/weather-go-api/model"
	"github.com/rs/zerolog/log"
)

// PlaceService is responsible for querying locations and regions
type PlaceService interface {
	GetLocations() ([]model.Location, error)
	GetRegions() ([]model.Region, error)
	SeedLocations([]model.Location) ([]model.Location, error)
	SeedRegions([]model.Region) ([]model.Region, error)
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
