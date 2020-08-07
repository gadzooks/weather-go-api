package main

import (
	"context"
	"fmt"
	"github.com/gadzooks/weather-go-api/middleware"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gadzooks/weather-go-api/config"
)

func main() {
	// show filename and line number in logs
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	r := config.NewRouter()
	config.AddAPISubRouterForPlaces(r)
	config.AddV2APISubRouterForPlaces(r)

	// connect to mongodb
	//FIXME get user, passwd, dbname from ENV

	// add middleware
	handler := middleware.WithCors(r)
	handler = middleware.WithResponseTimeLogging(handler)

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	mongoUser := os.Getenv("MONGO_USER")
	mongoPass := os.Getenv("MONGO_PWD")
	mongoDB := os.Getenv("MONGO_DB")
	uri := fmt.Sprintf(
		"mongodb+srv://%s:%s@weather-uqgvj.mongodb.net/%s?authSource=admin&replicaSet=Weather-shard-0&readPreference=primary",
		mongoUser,
		mongoPass,
		mongoDB,
	)
	log.Println("connecting via : ", uri)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("starting server at 8080")
	err = http.ListenAndServe(":8080", handler)
	if err != nil {
		log.Fatalf("error running server : %v", err)
	}

}
