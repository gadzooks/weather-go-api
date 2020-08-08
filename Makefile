# capture git sha & tag
export GIT_COMMIT_SHA = $(shell git rev-parse HEAD)
export GIT_COMMIT_TAG = $(shell git tag --points-at HEAD)

# gather project dependencies
dependencies:
	dep ensure

# update & gather project dependencies
dependencies-update:
	dep ensure -update

# generate swagger.json
swagger:
	#rm -f swagger/swagger-ui.yaml
	#swagger generate spec -o swagger/swagger-ui.yaml
	swagger validate swagger-ui.yaml

# view swagger documentation locally
swagger-view: swagger
	swagger serve swagger-ui.yaml

# build using prod targets
build-local:
	rm -rf ./bin
	env GOOS=darwin GOARCH=amd64 go build -o bin/dist/sample-go-api

# GO111MODULE=on go get github.com/golang/mock/mockgen@v1.4.3
mockgen:
	mockgen -source=controller/place_controller.go -destination=controller/place_controller_mock.go -package=controller
	mockgen -source=controller/v2/place_controller.go -destination=controller/v2/place_controller_mock.go -package=v2

build:
	mkdir -p bin
	rm -rf ./bin
	mkdir -p bin/dist
	env GOOS=linux GOARCH=amd64 go build -o bin/dist/sample-go-api

run:
	./bin/dist/sample-go-api

# run unit tests (all)
unit-test:
	rm -f coverage.*
	MONGO_DB_HOST=localhost \
	MONGO_DB_PORT=27017 \
	go test ./... -covermode=count -v -coverprofile coverage.out  && \
		go tool cover -html=coverage.out -o=coverage.html && \
		go tool cover -func=coverage.out
