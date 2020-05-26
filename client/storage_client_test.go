package client

import "testing"

func TestStorageClientImpl_QueryLocations(t *testing.T) {
	//Arrange
	client := NewStorageClient()

	//Act
	locations, err := client.QueryLocations("../data/")

	//Assert
	if err != nil {
		t.Errorf("could not load file : %v", err)
	}

	if len(locations) <= 0 {
		t.Errorf("could not parse locations. Found : %v", locations)
	}
}

func TestStorageClientImpl_QueryRegions(t *testing.T) {
	//Arrange
	client := NewStorageClient()

	//Act
	regions, err := client.QueryRegions("../data/")

	//Assert
	if err != nil {
		t.Errorf("could not load file : %v", err)
	}

	if len(regions) <= 0 {
		t.Errorf("could not parse regions. Found : %v", regions)
	}
}
