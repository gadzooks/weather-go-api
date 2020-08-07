package v2

import v1 "github.com/gadzooks/weather-go-api/client"

type LocationData = v1.LocationData
type RegionData = v1.RegionData

// StorageClient queries location db
type StorageClient interface {
	QueryLocations() (map[string]LocationData, error)
	QueryRegions() (map[string]RegionData, error)
	SeedLocations() (map[string]LocationData, error)
	SeedRegions() (map[string]RegionData, error)
}

// StorageClientImpl implements LocationClient interface
type StorageClientImpl struct {
	MongoUri        string
	Locations       map[string]LocationData
	locationsLoaded bool
	locationError   error
	Regions         map[string]RegionData
	regionsLoaded   bool
	regionError     error
}

func NewStorageClient(mongoUri string) StorageClient {
	return &StorageClientImpl{
		locationsLoaded: false,
		regionsLoaded:   false,
		MongoUri:        mongoUri,
	}
}

func (lci *StorageClientImpl) QueryRegions() (map[string]RegionData, error) {
	return nil, nil
}

func (lci *StorageClientImpl) QueryLocations() (map[string]LocationData, error) {
	return nil, nil
}

func (lci *StorageClientImpl) SeedRegions() (map[string]RegionData, error) {
	return nil, nil
}

func (lci *StorageClientImpl) SeedLocations() (map[string]LocationData, error) {
	return nil, nil
}
