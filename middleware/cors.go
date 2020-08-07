package middleware

import (
	"github.com/rs/cors"
	"net/http"
)

func WithCors(handler http.Handler) http.Handler {
	return cors.Default().Handler(handler)
}
