package client

import (
	"github.com/gadzooks/weather-go-api/model"
	"github.com/rs/zerolog/log"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// StorageClient queries location db
type StorageClient interface {
	//FIXME dataDir is not required
	QueryLocations(dataDir string) (map[string]model.Location, error)
	//FIXME dataDir is not required
	QueryRegions(dataDir string) (map[string]model.Region, error)
}

// StorageClientImpl implements LocationClient interface
type StorageClientImpl struct {
	DataDir         string
	Locations       map[string]model.Location
	locationsLoaded bool
	locationError   error
	Regions         map[string]model.Region
	regionsLoaded   bool
	regionError     error
}

func NewStorageClient(dataDir string) StorageClient {
	return &StorageClientImpl{
		locationsLoaded: false,
		regionsLoaded:   false,
		DataDir:         dataDir,
	}
}

func (lci *StorageClientImpl) QueryRegions(dataDir string) (map[string]model.Region, error) {
	if lci.regionsLoaded {
		return lci.Regions, lci.regionError
	}
	content, err := ioutil.ReadFile(lci.DataDir + "/regions.yml")
	if err != nil {
		lci.regionError = err
		log.Fatal().Msg(err.Error())
	}

	var results map[string]model.Region
	err = yaml.Unmarshal(content, &results)
	lci.regionError = err

	lci.regionsLoaded = true
	lci.Regions = results
	return results, err
}

func (lci *StorageClientImpl) QueryLocations(dataDir string) (map[string]model.Location, error) {
	if lci.locationsLoaded {
		return lci.Locations, lci.locationError
	}
	content, err := ioutil.ReadFile(lci.DataDir + "/locations.yml")
	if err != nil {
		lci.locationError = err
		return lci.Locations, lci.locationError
	}

	var results map[string]model.Location
	err = yaml.Unmarshal(content, &results)
	lci.locationError = err

	lci.locationsLoaded = true
	lci.Locations = results
	return results, err
}
