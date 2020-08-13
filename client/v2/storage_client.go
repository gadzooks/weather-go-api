package v2

import (
	"context"
	"fmt"
	"github.com/gadzooks/weather-go-api/model"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type Location = model.Location
type Region = model.Region

// StorageClient queries mongdb db
type StorageClient interface {
	QueryLocations() (map[string]Location, error)
	QueryRegions() (map[string]Region, error)
	CreateLocation(model.Location) (model.Location, error)
	CreateRegion(model.Region) (model.Region, error)
	DeleteAllRegions() error
	DeleteAllLocations() error
	FindRegion(string) (model.Region, error)
	UpdateRegion(map[string]string) error
	DeleteRegion(string) error
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

func (lci *StorageClientImpl) UpdateRegion(params map[string]string) error {

	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		return fmt.Errorf("invalid id '%s' provided. %w", params["id"], InvalidInputError)
	}

	ctx, cancel := lci.getContext()
	defer cancel()
	collection := lci.MongoClient.Database(lci.dbName).Collection(regionCollection)

	var updateSubQuery bson.D
	description := params["description"]
	if len(description) > 0 {
		updateSubQuery = bson.D{{"description", description}}
	}

	updateResult, err := collection.UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.D{
			{"$set", updateSubQuery},
		},
	)
	log.Info().Msgf("updated %d region", updateResult.ModifiedCount)
	return err
}

// from https://blog.golang.org/go1.13-errors
// allows callers to do something like this :
// if err := pkg.DoSomething(); errors.Is(err, pkg.InvalidInputError) { ... }
var InvalidInputError = errors.New("InvalidInputError")
var NotFoundError = errors.New("NotFoundError")

func (lci *StorageClientImpl) DeleteRegion(id string) error {
	oID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("invalid id '%s' provided. %w", id, InvalidInputError)
	}

	query := bson.M{"_id": oID}
	ctx, cancel := lci.getContext()
	defer cancel()
	collection := lci.MongoClient.Database(lci.dbName).Collection(regionCollection)

	log.Info().Msgf("executing query : %v", query)
	result, err := collection.DeleteOne(ctx, query)
	if err != nil {
		return err
	}
	log.Info().Msgf("DeleteOne removed %v document(s)\n", result.DeletedCount)
	return nil
}

func (lci *StorageClientImpl) FindRegion(id string) (model.Region, error) {
	oID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return model.Region{}, fmt.Errorf("invalid id '%s' provided. %w", id, InvalidInputError)
	}

	query := bson.M{"_id": oID}

	log.Info().Msgf("executing query : %v", query)
	ctx, cancel := lci.getContext()
	defer cancel()
	collection := lci.MongoClient.Database(lci.dbName).Collection(regionCollection)
	var region model.Region
	err = collection.FindOne(ctx, query).Decode(&region)
	if err != nil {
		return model.Region{}, fmt.Errorf("region not found for '%s': %w", id, NotFoundError)
	}
	return region, nil
}

func (lci *StorageClientImpl) DeleteAllRegions() error {
	ctx, cancel := lci.getContext()
	defer cancel()
	collection := lci.MongoClient.Database(lci.dbName).Collection(regionCollection)
	count, _ := collection.CountDocuments(ctx, nil)
	log.Info().Msgf("deleting %d regions", count)
	return collection.Drop(ctx)
}

func (lci *StorageClientImpl) DeleteAllLocations() error {
	ctx, cancel := lci.getContext()
	defer cancel()
	collection := lci.MongoClient.Database(lci.dbName).Collection(locationCollection)
	count, _ := collection.CountDocuments(ctx, nil)
	log.Info().Msgf("deleting %d locations", count)
	return collection.Drop(ctx)
}

func (lci *StorageClientImpl) QueryRegions() (map[string]Region, error) {
	ctx, cancel := lci.getContext()
	defer cancel()
	collection := lci.MongoClient.Database(lci.dbName).Collection(regionCollection)

	var results = make(map[string]Region)
	cursor, err := collection.Find(ctx, bson.M{})
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
	cursor, err := collection.Find(ctx, bson.M{})
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

func (lci *StorageClientImpl) getContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 5*time.Second)
}
