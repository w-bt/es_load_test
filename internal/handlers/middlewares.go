package handlers

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

const (
	requestIDKey          = "X-Request-ID"
	frameOptionsKey       = "X-Frame-Options"
	contentTypeKey        = "Content-Type"
	transferEncodingKey   = "Transfer-Encoding"
	xssProtectionKey      = "X-XSS-Protection"
	contentTypeOptionsKey = "X-Content-Type-Options"
	cacheControlKey       = "Cache-Control"
)

func withDefaultResponseHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(wr http.ResponseWriter, req *http.Request) {
		wr.Header().Add(frameOptionsKey, "SAMEORIGIN")
		wr.Header().Add(contentTypeKey, "application/json; charset=utf-8")
		wr.Header().Add(transferEncodingKey, "chunked")
		wr.Header().Add(xssProtectionKey, "1; mode=block")
		wr.Header().Add(contentTypeOptionsKey, "nosniff")
		wr.Header().Add(cacheControlKey, "max-age=0, private, must-revalidate")

		rid := req.Header.Get(requestIDKey)
		if rid == "" {
			rid = uuid.New().String()
			req.Header.Set(requestIDKey, rid)
		}

		wr.Header().Add(requestIDKey, rid)
		req = req.WithContext(context.WithValue(req.Context(), requestIDKey, rid))
		next.ServeHTTP(wr, req)
	})
}
