# Weather Config REST API - Go REST API

Sample Go api to learn Go

## dirs

- controller : REST controller
- model : Data Transfer Objects
- domain : Provide functionality based on DTO types
- service : Business logic lives here

## build

go build main.go from root dir

## goals of this service
- create sample CRUD REST api in Go which has
    - well defined classes for model, domain, controller
    - Makefile
    - deploy with kubernetes, along side rails app
    - logging, error handling
    - log storage options
    - connecting to mongodb to get props   
    - JWT auth ? for service to service authentication
    - use different envs : dev, staging, prod
    - unit tests
    - postman tests
    - newman tests
    - swagger
    
    
## order of implementation : 
- CRUD operations for location config, in memory
- store CRUD in mongodb
- 
