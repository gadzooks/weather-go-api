package v2

import (
	"context"
	v1 "github.com/gadzooks/weather-go-api/client"
	"github.com/gadzooks/weather-go-api/model"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type LocationData = v1.LocationData
type RegionData = v1.RegionData

// StorageClient queries location db
type StorageClient interface {
	QueryLocations() (map[string]LocationData, error)
	QueryRegions() (map[string]RegionData, error)
	CreateLocation(model.Location) (map[string]LocationData, error)
	CreateRegion(model.Region) (model.Region, error)
}

// StorageClientImpl implements LocationClient interface
type StorageClientImpl struct {
	MongoClient     *mongo.Client
	Locations       map[string]LocationData
	locationsLoaded bool
	locationError   error
	Regions         map[string]RegionData
	regionsLoaded   bool
	regionError     error
}

func NewStorageClient(mongoClient *mongo.Client) StorageClient {
	return &StorageClientImpl{
		locationsLoaded: false,
		regionsLoaded:   false,
		MongoClient:     mongoClient,
	}
}

func (lci *StorageClientImpl) QueryRegions() (map[string]RegionData, error) {
	return nil, nil
}

func (lci *StorageClientImpl) QueryLocations() (map[string]LocationData, error) {
	return nil, nil
}

func (lci *StorageClientImpl) CreateRegion(data model.Region) (model.Region, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := lci.MongoClient.Database("test").Collection("regions")

	_, err := collection.InsertOne(ctx, data)
	if err != nil {
		return data, err
	}
	// FIXME : pass id back, check for existing value and upsert
	// data.ID = res.InsertedID
	return data, err
}

func (lci *StorageClientImpl) CreateLocation(model.Location) (map[string]LocationData, error) {
	return nil, nil
}
