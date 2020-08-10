# Weather Config REST API - Go REST API

Sample REST Go api to get locations and regions to be used by Ruby on Rails weather website.

## Goals of this service
Create sample CRUD REST api in Go which has :
- Follow REST naming conventions
- Use swagger for documentation
- Well defined interfaces for model, domain, controller
- All commands can be run via Makefile targets
- Deploy service and swagger docs via docker
- Code coverage, unit and integration tests

## Design Patterns / Best practices

### Use `docker-compose` to build and run `go service` and `swagger`
- https://github.com/gadzooks/weather-go-api/blob/master/docker-compose.yml
- https://github.com/gadzooks/weather-go-api/blob/master/Dockerfile

#### Use middleware to handle tasks like CORS, logging request performance etc
```go
//ServeHTTP handles the request by passing it to the real
//handler and logging the request details
func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	l.handler.ServeHTTP(w, r)
	log.Info().Msgf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
}
```

### Use middleware to add unique reqId per request to context
```shell script
6:07AM INF service/v2/place_service.go:41 > insert location : {ObjectID("000000000000000000000000") north bend snowqualmie_region North Bend 47.497428 -121.786648 db086e5e85941a02ae188f726f7e9e2c}
 reqId=74590ef0-2c7b-4e62-bfef-03ae13fa85a3
6:07AM INF service/v2/place_service.go:47 > inserted location with id : 5f30e4046c7600d75072329d reqId=74590ef0-2c7b-4e62-bfef-03ae13fa85a3
6:07AM INF middleware/request_logger.go:19 > POST /v2/locations 3.15215768s reqId=74590ef0-2c7b-4e62-bfef-03ae13fa85a3
```

#### Follow REST naming convention and swagger documentation style
```go
	// FindLocations swagger:route GET /locations locations findLocations
	//
	// Finds a location set
	//
	// Consumes:
	// - application/json
	//
	// Produces:
	// - application/json
	//
	// Responses:
	// 200: []location
	api.HandleFunc("/locations", placesCtrl.FindLocations).Methods("GET")
```

#### Rest controller(s) expose endpoints via PlaceController `interface`
```go
type PlaceController interface {
	FindLocations(w http.ResponseWriter, r *http.Request)
	FindRegions(w http.ResponseWriter, r *http.Request)
}
```

#### Service `package` handles business logic via PlaceService `interface`
```go
// PlaceService is responsible for querying locations and regions
type PlaceService interface {
	GetLocations() ([]model.Location, error)
	GetRegions() ([]model.Region, error)
}

// NewPlaceServiceImpl implements NewPlaceService
type NewPlaceServiceImpl struct {
	client client.StorageClient
}
```

#### Client `package` handles external clients
```go
// StorageClient queries location db
type StorageClient interface {
	QueryLocations(dataDir string) (map[string]model.Location, error)
	QueryRegions(dataDir string) (map[string]model.Region, error)
}
```

#### Model `pacakge` contains Data Transfer Objects (DTOs)
```go
/*
snowqualmie_region:
  search_key: '04d37e830680c65b61df474e7e655d64'
  description: 'Snowqualmie Region'
*/
type Region struct {
	Name        string `json:"name"`
	SearchKey   string `json:"searchKey"`
	Description string `json:"description"`
}
```

#### Testing
- we use gomock and mocken
```shell script
# to install
go get github.com/golang/mock/gomock
go get github.com/golang/mock/mockgen
mockgen -source=controller/place_controller.go -destination=controller/place_controller_mock.go -package=controller
```
- example integration test : 
```go
func TestGetLocations(t *testing.T) {
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

```

#### we use swagger to generate living documentation
https://goswagger.io/

#### build / generate mocks / run tests is via `Makefile` targets
https://github.com/gadzooks/weather-go-api/blob/master/Makefile

#### dependencies
we use `dep` to mange dependencies

#### code coverage via `go test`
https://github.com/gadzooks/weather-go-api/blob/master/coverage.out

#### Todo
- postman and newman tests
- deploy with kubernetes, alongside other apps
- use different envs : dev, staging, prod
- JWT auth ? for service to service authentication
- connecting to mongodb to get props   
