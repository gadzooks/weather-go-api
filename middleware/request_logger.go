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
	log.Info().Msgf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
}

//WithResponseTimeLogging constructs a new RequestLogger middleware handler
func WithResponseTimeLogging(handlerToWrap http.Handler) *RequestLogger {
	return &RequestLogger{handlerToWrap}
}
