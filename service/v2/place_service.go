package v2

import (
	v2Client "github.com/gadzooks/weather-go-api/client/v2"
	"github.com/gadzooks/weather-go-api/model"
	"log"
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
	return nil, nil
}

func (r PlaceServiceImpl) SeedLocations(data []model.Location) ([]model.Location, error) {
	// loc, err := r.client.CreateLocation(data)

	return nil, nil
}

func (r PlaceServiceImpl) GetRegions() ([]model.Region, error) {
	return nil, nil
}

func (r PlaceServiceImpl) SeedRegions(data []model.Region) ([]model.Region, error) {
	// collection := client.Database("testing").Collection("numbers")
	var inserted []model.Region
	for _, region := range data {
		log.Printf("insert region : %v\n", region)
		r, err := r.client.CreateRegion(region)
		if err != nil {
			log.Printf("error inserting region : %v\n", err)
		} else {
			inserted = append(inserted, region)
			log.Printf("inserted region with id : %v", r.ID.Hex())
		}
	}
	return inserted, nil
}

// NewPlaceService creates new instance of PlaceService
func NewPlaceService(client v2Client.StorageClient) PlaceService {
	return &PlaceServiceImpl{client: client}
}
