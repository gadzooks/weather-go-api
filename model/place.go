package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*
Mongodb validation :
{
  $jsonSchema: {
    bsonType: 'object',
    required: [
      'name',
      'region',
      'latitude',
      'longitude',
      'subRegion'
    ],
    properties: {
      _id: {},
      name: {
        bsonType: 'string',
        description: '\'name\' is required and should be a string'
      },
      region: {
        bsonType: 'string',
        description: '\'region\' is required and should be a string'
      },
      subRegion: {
        bsonType: 'string',
        description: '\'subRegion\' is required and should be a string'
      },
      latitude: {
        bsonType: 'double',
        description: '\'latitude\' is required and should be a valid double'
      },
      longitude: {
        bsonType: 'double',
        description: '\'longitude\' is required and should be a valid double'
      }
    }
  }
}

gold bar:
  name: 'gold bar'
  region: central_cascades
  description: 'Stevens Pass - West'
  latitude: 47.8090
  longitude: -121.5738
  sub_region: '637634387ca38685f89162475c7fc1d2'
*/
type Location struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	Name        string             `json:"name" yaml:"name" bson:"name"`
	Region      string             `json:"region" yaml:"region" bson:"region"`
	Description string             `json:"description" yaml:"description" bson:"description"`
	Latitude    float64            `json:"latitude" yaml:"latitude" bson:"latitude"`
	Longitude   float64            `json:"longitude" yaml:"longitude" bson:"longitude"`
	SubRegion   string             `json:"subRegion" yaml:"sub_region" bson:"subRegion"`
}

/*
mongodb validation :
{
  $jsonSchema: {
    bsonType: 'object',
    required: [
      'name',
      'searchKey',
      'description'
    ],
    properties: {
      _id: {},
      name: {
        bsonType: 'string',
        description: '\'name\' is required and should be a string'
      },
      searchKey: {
        bsonType: 'string',
        description: '\'searchKey\' is required and should be a string'
      },
      description: {
        bsonType: 'string',
        description: '\'description\' is required and should be a string'
      }
    }
  }
}

snowqualmie_region:
  search_key: '04d37e830680c65b61df474e7e655d64'
  description: 'Snowqualmie Region'
*/
type Region struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	Name        string             `json:"name" yaml:"name" bson:"name"`
	SearchKey   string             `json:"searchKey" yaml:"search_key" bson:"searchKey"`
	Description string             `json:"description" yaml:"description" bson:"description"`
}
