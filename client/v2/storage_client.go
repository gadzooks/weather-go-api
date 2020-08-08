package v2

import (
	"context"
	"github.com/gadzooks/weather-go-api/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

type Location = model.Location
type Region = model.Region

// StorageClient queries location db
type StorageClient interface {
	QueryLocations() (map[string]Location, error)
	QueryRegions() (map[string]Region, error)
	CreateLocation(model.Location) (model.Location, error)
	CreateRegion(model.Region) (model.Region, error)
	DeleteAllRegions() error
	DeleteAllLocations() error
}

// StorageClientImpl implements LocationClient interface
type StorageClientImpl struct {
	MongoClient *mongo.Client
	Locations   map[string]Location
	Regions     map[string]Region
	dbName      string
}

func NewStorageClient(mongoClient *mongo.Client) StorageClient {
	return &StorageClientImpl{
		MongoClient: mongoClient,
		dbName:      "test", //FIXME this should come from ENV
	}
}

const regionCollection = "regions"
const locationCollection = "locations"

func (lci *StorageClientImpl) getContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 5*time.Second)
}

func (lci *StorageClientImpl) DeleteAllRegions() error {
	ctx, cancel := lci.getContext()
	defer cancel()
	collection := lci.MongoClient.Database(lci.dbName).Collection(regionCollection)
	count, _ := collection.CountDocuments(ctx, nil)
	log.Printf("deleting %d regions", count)
	return collection.Drop(ctx)
}

func (lci *StorageClientImpl) DeleteAllLocations() error {
	ctx, cancel := lci.getContext()
	defer cancel()
	collection := lci.MongoClient.Database(lci.dbName).Collection(locationCollection)
	count, _ := collection.CountDocuments(ctx, nil)
	log.Printf("deleting %d locations", count)
	return collection.Drop(ctx)
}

func (lci *StorageClientImpl) QueryRegions() (map[string]Region, error) {
	ctx, cancel := lci.getContext()
	defer cancel()
	collection := lci.MongoClient.Database(lci.dbName).Collection(regionCollection)

	var results = make(map[string]Region)
	cursor, err := collection.Find(ctx, nil)
	if err != nil {
		return results, err
	}
	var regions []Region
	if err = cursor.All(ctx, &regions); err != nil {
		return results, err
	}
	for _, r := range regions {
		n := r.Name
		results[n] = r
	}

	return results, nil
}

func (lci *StorageClientImpl) QueryLocations() (map[string]Location, error) {
	ctx, cancel := lci.getContext()
	defer cancel()
	collection := lci.MongoClient.Database(lci.dbName).Collection(locationCollection)

	var results = make(map[string]Location)
	cursor, err := collection.Find(ctx, nil)
	if err != nil {
		return results, err
	}
	var locations []Location
	if err = cursor.All(ctx, &locations); err != nil {
		return results, err
	}
	for _, r := range locations {
		results[r.Name] = r
	}

	return results, nil
}

func (lci *StorageClientImpl) CreateRegion(data model.Region) (model.Region, error) {
	ctx, cancel := lci.getContext()
	defer cancel()

	collection := lci.MongoClient.Database(lci.dbName).Collection(regionCollection)

	data.ID = primitive.NewObjectID()
	_, err := collection.InsertOne(ctx, data)
	if err != nil {
		return data, err
	}
	return data, err
}

func (lci *StorageClientImpl) CreateLocation(data model.Location) (model.Location, error) {
	ctx, cancel := lci.getContext()
	defer cancel()
	collection := lci.MongoClient.Database(lci.dbName).Collection(locationCollection)

	data.ID = primitive.NewObjectID()
	_, err := collection.InsertOne(ctx, data)
	if err != nil {
		return data, err
	}
	return data, err
}
