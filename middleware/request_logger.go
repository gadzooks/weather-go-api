package middleware

import (
	"github.com/rs/zerolog/log"
	"net/http"
	"time"
)

//RequestLogger is a middleware handler that does request logging
type RequestLogger struct {
	handler http.Handler
}

//ServeHTTP handles the request by passing it to the real
//handler and logging the request details
func (l *RequestLogger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	l.handler.ServeHTTP(w, r)
	log.Logger = log.With().Str("reqTime", time.Since(start).String()).Logger()
	log.Info().Msgf("%s %s", r.Method, r.URL.Path)
}

//WithResponseTimeLogging constructs a new RequestLogger middleware handler
func WithResponseTimeLogging(handlerToWrap http.Handler) *RequestLogger {
	return &RequestLogger{handlerToWrap}
}
