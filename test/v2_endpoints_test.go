package test

import (
	"encoding/json"
	client "github.com/gadzooks/weather-go-api/client/v2"
	controller "github.com/gadzooks/weather-go-api/controller/v2"
	"github.com/gadzooks/weather-go-api/model"
	service "github.com/gadzooks/weather-go-api/service/v2"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetLocationsV2(t *testing.T) {
	req, err := http.NewRequest("GET", "/locations", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	storageClient := client.NewStorageClient("../data")
	svc := service.NewPlaceService(storageClient)
	placesCtrl := controller.NewPlaceController(svc)

	handler := http.HandlerFunc(placesCtrl.FindLocations)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Fatalf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	t.Skip("pending implementation")
	// Check the response body is what we expect.
	expected := 26
	var result *[]model.Location
	err = json.Unmarshal(rr.Body.Bytes(), &result)
	if err != nil {
		t.Fatalf("error unmarshaling results : %v", err)
	}
	if len(*result) != expected {
		t.Fatalf("got %v want %v",
			len(*result), expected)
	}
}

func TestGetRegionsV2(t *testing.T) {
	req, err := http.NewRequest("GET", "/regions", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	storageClient := client.NewStorageClient("../data")
	svc := service.NewPlaceService(storageClient)
	placesCtrl := controller.NewPlaceController(svc)

	handler := http.HandlerFunc(placesCtrl.FindRegions)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Fatalf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	t.Skip("pending implementation")
	// Check the response body is what we expect.
	expected := 7
	var result *[]model.Region
	err = json.Unmarshal(rr.Body.Bytes(), &result)
	if err != nil {
		t.Fatalf("error unmarshaling results : %v", err)
	}
	if len(*result) != expected {
		t.Fatalf("got %v want %v",
			len(*result), expected)
	}
}
