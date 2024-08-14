package utils

import (
	"log"
	"net/http"
)

// Simple example piece of middleware that logs each request
func RequestLoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
