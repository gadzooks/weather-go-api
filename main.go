package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/gadzooks/weather-go-api/middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gadzooks/weather-go-api/config"
)

const MongoTimeOutInSeconds = 10

func main() {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	// UNIX Time is faster and smaller than most timestamps
	// If you set zerolog.TimeFieldFormat to an empty string,
	// logs will write with UNIX time
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	// show filename and line number in logs
	log.Logger = log.With().Caller().Logger()
	//log.Logger = log.With().Str("requestId", "bar").Logger()
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	mongoClient := mongoConnect(ctx)

	router := config.NewRouter()
	config.AddAPISubRouterForPlaces(router)
	config.AddV2APISubRouterForPlaces(router, mongoClient)

	// add middleware
	handler := middleware.WithCors(router)
	handler = middleware.WithResponseTimeLogging(handler)
	handler = middleware.Middleware(handler)

	srv := &http.Server{
		Handler: handler,
		Addr:    "127.0.0.1:8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		log.Info().Msg("starting server at 8080")
		if err := srv.ListenAndServe(); err != nil {
			log.Error().Msg(err.Error())
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	// ctx, cancel = context.WithTimeout(context.Background(), wait)
	// defer cancel()

	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Info().Msg("shutting down")
	os.Exit(0)
}

func mongoConnect(ctx context.Context) *mongo.Client {
	mongoUser := os.Getenv("MONGO_USER")
	mongoPass := os.Getenv("MONGO_PWD")
	mongoDB := os.Getenv("MONGO_DB")
	uri := fmt.Sprintf(
		"mongodb+srv://%s:%s@weather-uqgvj.mongodb.net/%s?authSource=admin&replicaSet=Weather-shard-0&readPreference=primary",
		mongoUser,
		mongoPass,
		mongoDB,
	)
	log.Info().Msgf("connecting via : ", uri)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
	log.Info().Msg("connected to mongo successfully !")

	return client
}
