# Start from golang base image
FROM golang:alpine as builder

ENV GO111MODULE=on

# Add Maintainer info
LABEL maintainer="Amit Karwande <karwande@gmail.com>"

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

# Set the current working directory inside the container
WORKDIR /go/src/github.com/gadzooks/weather-go-api

# Copy go mod and sum files
# This way if dependencies do not change then we wont re-download dependencies
# this will make the build process fast
# Later we copy all the files so that we get any code changes. Since the deps were
# already accounted for in the earlier step, we can go ahead and just rebuild the service
COPY go.mod go.sum Makefile data ./

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the working Directory inside the container
COPY . .

# Build the Go app
RUN GOOS=linux GOARCH=amd64 go build -o weather-api

############################
# STEP 2 build a small image
############################
# Start a new stage from scratch
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage. Observe we also copied the .env file
COPY --from=builder /go/src/github.com/gadzooks/weather-go-api/weather-api .
COPY --from=builder /go/src/github.com/gadzooks/weather-go-api/data ./data/
COPY --from=builder /go/src/github.com/gadzooks/weather-go-api/.env .

# Expose port 8080 to the outside world
EXPOSE 8080

#Command to run the executable
CMD ["./weather-api"]