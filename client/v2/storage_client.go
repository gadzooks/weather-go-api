package v2

import (
	v1 "github.com/gadzooks/weather-go-api/client"
	"github.com/gadzooks/weather-go-api/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type LocationData = v1.LocationData
type RegionData = v1.RegionData

// StorageClient queries location db
type StorageClient interface {
	QueryLocations() (map[string]LocationData, error)
	QueryRegions() (map[string]RegionData, error)
	SeedLocations([]model.Location) (map[string]LocationData, error)
	SeedRegions([]model.Region) (map[string]RegionData, error)
}

// StorageClientImpl implements LocationClient interface
type StorageClientImpl struct {
	MongoUri             string
	mongoConnectionError error
	mongoClient          *mongo.Client
	Locations            map[string]LocationData
	locationsLoaded      bool
	locationError        error
	Regions              map[string]RegionData
	regionsLoaded        bool
	regionError          error
}

func NewStorageClient(mongoUri string) StorageClient {
	return &StorageClientImpl{
		locationsLoaded:      false,
		regionsLoaded:        false,
		MongoUri:             mongoUri,
		mongoClient:          nil,
		mongoConnectionError: nil,
	}
}

func (lci *StorageClientImpl) QueryRegions() (map[string]RegionData, error) {
	return nil, nil
}

func (lci *StorageClientImpl) QueryLocations() (map[string]LocationData, error) {
	return nil, nil
}

func (lci *StorageClientImpl) SeedRegions([]model.Region) (map[string]RegionData, error) {
	return nil, nil
}

func (lci *StorageClientImpl) SeedLocations([]model.Location) (map[string]LocationData, error) {
	return nil, nil
}
