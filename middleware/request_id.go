package middleware

import (
	"context"
	"github.com/google/uuid"
	"net/http"
)

type ContextKey string

const (
	// ContextKeyReqID is the context key for RequestID
	ContextKeyReqID ContextKey = "requestID"

	// HTTPHeaderNameRequestID has the name of the header for request ID
	HTTPHeaderNameRequestID = "X-Request-ID"
)

// GetReqID will get reqID from a http request and return it as a string
func GetReqID(ctx context.Context) string {
	reqID := ctx.Value(ContextKeyReqID)
	if ret, ok := reqID.(string); ok {
		return ret
	}
	return ""
}

// attachReqID will attach a brand new request ID to a http request
func attachReqID(ctx context.Context) context.Context {
	reqID := uuid.New()
	return context.WithValue(ctx, ContextKeyReqID, reqID.String())
}

// WithRequestIdInLogger will attach the reqID to the http.Request and add reqID to http header in the response
func WithRequestIdInLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := attachReqID(r.Context())
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
		h := w.Header()
		h.Add(HTTPHeaderNameRequestID, GetReqID(ctx))
	})
}
