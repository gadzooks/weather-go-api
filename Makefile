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
	rm -f ./.docs/swagger.json
	go generate
	swagger validate ./.docs/swagger.json

# view swagger documentation locally
swagger-view: swagger
	swagger serve ./.docs/swagger.json

# build using prod targets
build:
	rm -rf ./bin
	env GOOS=linux GOARCH=amd64 go build -o bin/dist/sample-go-api

# run unit tests (all)
unit-test:
	rm -f coverage.*
	MONGO_DB_HOST=localhost \
	MONGO_DB_PORT=27017 \
	go test ./... -covermode=count -v -coverprofile coverage.out  && \
		go tool cover -html=coverage.out -o=coverage.html && \
		go tool cover -func=coverage.out
