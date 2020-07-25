package service

import (
	"github.com/gadzooks/weather-api-go/client"
	"github.com/gadzooks/weather-api-go/model"
)

// PlaceService is responsible for querying locations
type PlaceService interface {
	GetLocations() ([]model.Location, error)
	GetRegions() ([]model.Region, error)
}

// NewPlaceServiceImpl implements NewPlaceService
type NewPlaceServiceImpl struct {
	client client.StorageClient
}

const dataDir = "data/"

func (r NewPlaceServiceImpl) GetLocations() ([]model.Location, error) {
	loc, err := r.client.QueryLocations(dataDir)

	if err != nil {
		return nil, err
	}

	results := make([]model.Location, 0, len(loc))
	for _, v := range loc {
		results = append(results, model.Location(v))
	}

	return results, nil
}

func (r NewPlaceServiceImpl) GetRegions() ([]model.Region, error) {
	loc, err := r.client.QueryRegions(dataDir)

	if err != nil {
		return nil, err
	}

	results := make([]model.Region, 0, len(loc))
	for k, v := range loc {
		results = append(results, model.Region{Name: k, SearchKey: v.SearchKey, Description: v.Description})
	}

	return results, nil
}

// NewPlaceService creates new instance of PlaceService
func NewPlaceService(client client.StorageClient) PlaceService {
	return &NewPlaceServiceImpl{client: client}
}
