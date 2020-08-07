package v2

import (
	v2Client "github.com/gadzooks/weather-go-api/client/v2"
	"github.com/gadzooks/weather-go-api/model"
)

// PlaceService is responsible for querying locations and regions
type PlaceService interface {
	GetLocations() ([]model.Location, error)
	GetRegions() ([]model.Region, error)
	SeedLocations() ([]model.Location, error)
	SeedRegions() ([]model.Region, error)
}

// PlaceServiceImpl implements PlaceService
type PlaceServiceImpl struct {
	client v2Client.StorageClient
}

func (r PlaceServiceImpl) GetLocations() ([]model.Location, error) {
	return nil, nil
}

func (r PlaceServiceImpl) SeedLocations() ([]model.Location, error) {
	return nil, nil
}

func (r PlaceServiceImpl) GetRegions() ([]model.Region, error) {
	return nil, nil
}

func (r PlaceServiceImpl) SeedRegions() ([]model.Region, error) {
	return nil, nil
}

// NewPlaceService creates new instance of PlaceService
func NewPlaceService(client v2Client.StorageClient) PlaceService {
	return &PlaceServiceImpl{client: client}
}
