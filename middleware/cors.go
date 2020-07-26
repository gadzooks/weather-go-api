package middleware

import (
	"github.com/rs/cors"
	"net/http"
)

func SetupGlobalMiddleware(handler http.Handler) http.Handler {
	handleCORS := cors.Default().Handler

	return handleCORS(handler)
}
