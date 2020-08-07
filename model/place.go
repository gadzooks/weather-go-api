package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Location struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	Name        string             `json:"name" yaml:"name" bson:"name"`
	Region      string             `json:"region" yaml:"region" bson:"region"`
	Description string             `json:"description" yaml:"description" bson:"description"`
	Latitude    float64            `json:"latitude" yaml:"latitude" bson:"latitude"`
	Longitude   float64            `json:"longitude" yaml:"longitude" bson:"longitude"`
	SubRegion   string             `json:"subRegion" yaml:"sub_region" bson:"subRegion"`
}

/*
snowqualmie_region:
  search_key: '04d37e830680c65b61df474e7e655d64'
  description: 'Snowqualmie Region'
*/
type Region struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	Name        string             `json:"name" yaml:"name" bson:"name"`
	SearchKey   string             `json:"searchKey" yaml:"search_key" bson:"searchKey"`
	Description string             `json:"description" yaml:"description" bson:"description"`
}
