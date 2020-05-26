package client

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// StorageClient queries location db
type StorageClient interface {
	QueryLocations(dataDir string) (map[string]LocationData, error)
	QueryRegions(dataDir string) (map[string]RegionData, error)
}

// StorageClientImpl implements LocationClient interface
type StorageClientImpl struct {
	Locations       map[string]LocationData
	locationsLoaded bool
	locationError   error
	Regions         map[string]RegionData
	regionsLoaded   bool
	regionError     error
}

/*
gold bar:
  name: 'gold bar'
  region: central_cascades
  description: 'Stevens Pass - West'
  latitude: 47.8090
  longitude: -121.5738
  sub_region: '637634387ca38685f89162475c7fc1d2'
*/
// LocationData data returned by storage db service.
type LocationData struct {
	Name        string  `yaml:"name"`
	Region      string  `yaml:"region"`
	Description string  `yaml:"description"`
	Latitude    float64 `yaml:"latitude"`
	Longitude   float64 `yaml:"longitude"`
	SubRegion   string  `yaml:"sub_region"`
}

/*
snowqualmie_region:
  search_key: '04d37e830680c65b61df474e7e655d64'
  description: 'Snowqualmie Region'
*/
// RegionData is data returned by storage db service
type RegionData struct {
	SearchKey   string `yaml:"search_key"`
	Description string `yaml:"description"`
}

func NewStorageClient() StorageClient {
	return &StorageClientImpl{
		locationsLoaded: false,
		regionsLoaded:   false,
	}
}

func (lci *StorageClientImpl) QueryRegions(dataDir string) (map[string]RegionData, error) {
	if lci.regionsLoaded {
		return lci.Regions, lci.locationError
	}
	content, err := ioutil.ReadFile(dataDir + "/regions.yml")
	if err != nil {
		lci.regionError = err
		log.Fatal(err)
	}

	var results map[string]RegionData
	err = yaml.Unmarshal(content, &results)
	lci.regionError = err

	lci.regionsLoaded = true
	lci.Regions = results
	return results, err
}

func (lci *StorageClientImpl) QueryLocations(dataDir string) (map[string]LocationData, error) {
	if lci.locationsLoaded {
		return lci.Locations, lci.regionError
	}
	content, err := ioutil.ReadFile(dataDir + "/locations.yml")
	if err != nil {
		log.Fatal(err)
		lci.locationError = err
	}

	var results map[string]LocationData
	err = yaml.Unmarshal(content, &results)
	lci.locationError = err

	lci.locationsLoaded = true
	lci.Locations = results
	return results, err
}
