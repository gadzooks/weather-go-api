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
	handler = middleware.WithRequestIdInLogger(handler)

	bindIp, found := os.LookupEnv("WEATHER_BINDING_IP")
	if !found {
		bindIp = "127.0.0.1"
	}

	bindPort, found := os.LookupEnv("WEATHER_BINDING_PORT")
	if !found {
		bindPort = "8080"
	}

	addr := fmt.Sprintf("%s:%s", bindIp, bindPort)

	srv := &http.Server{
		Handler: handler,
		Addr:    addr,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		log.Info().Msgf("starting server at %s", addr)
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

func buildMongoUri() string {
	uri, found := os.LookupEnv("MONGO_CONNECTION_STRING")
	if found {
		log.Debug().Msgf("using mongo connection string : %s", uri)
		return uri
	}

	mongoConnectPrefix, found := os.LookupEnv("MONGO_PREFIX")
	if !found {
		mongoConnectPrefix = "mongodb+srv"
	}

	mongoUser, _ := os.LookupEnv("MONGO_USER") //weather-dev
	if !found {
		mongoUser = "weather-dev"
	}
	mongoPass := os.Getenv("MONGO_PWD")
	mongoDB, found := os.LookupEnv("MONGO_DB")
	if !found {
		mongoDB = "weatherDevDb"
	}
	mongoUri, found := os.LookupEnv("MONGO_SERVER_URI")
	if !found {
		mongoUri = "weather-uqgvj.mongodb.net"
	}

	uri = fmt.Sprintf(
		"%s://%s:%s@%s/%s?authSource=admin&readPreference=primary",
		mongoConnectPrefix,
		mongoUser,
		mongoPass,
		mongoUri,
		mongoDB,
	)

	//uri = "mongodb://integUser:integPass@127.0.0.1:27017/test"
	//uri = "mongodb://integUser:integPass@mongodb:27017/test"
	return uri
}

func mongoConnect(ctx context.Context) *mongo.Client {
	uri := buildMongoUri()
	log.Info().Msgf("connecting to mongo uri : %s", uri)
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
