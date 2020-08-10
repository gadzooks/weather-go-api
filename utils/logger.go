package utils

import (
	"context"
	"github.com/gadzooks/weather-go-api/middleware"
	"github.com/rs/zerolog/log"
)

func SetLoggerWithRequestId(ctx context.Context) {
	ctxRqId := middleware.GetReqID(ctx)
	if ctxRqId != "" {
		log.Logger = log.With().Str("reqId", ctxRqId).Logger()
	}
}
