package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/ztolley/goapi/internal/auth"
	"github.com/ztolley/goapi/internal/docs"
	"github.com/ztolley/goapi/internal/user"
	"github.com/ztolley/goapi/internal/utils"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() {
	// Use the new server router added in Go 1.22
	router := http.NewServeMux()

	// Create a new handler service for user related requests, pass it a reference
	// to the user store so it can make database repository calls then register the
	// routes it will handle
	userHandler := user.NewHandler(user.NewStore(s.db))
	userHandler.RegisterRoutes(router)

	// Create a handler for OpenAPI documentation and Swagger UI
	docs.RegisterRoutes(router)

	// Define the routes to exclude from JWT authentication
	// and create a new JWT authentication middleware with the excluded routes
	excludedRoutes := []string{
		"/openapi.yaml",
		"/openapi.json",
		"/swagger/",
		"/favicon.ico",
	}

	authMiddleware := auth.WithJWTAuth(excludedRoutes)

	// Define the middleware chain for all requests to include logging and authentication
	middlewareChain := MiddlewareChain(utils.RequestLoggerMiddleware, authMiddleware)

	// Setup the web server, this is a mixture of routes, middleare and the
	// address to listen on for requests
	server := http.Server{
		Addr:    s.addr,
		Handler: middlewareChain(router),
	}

	log.Printf("API server is running on %s", s.addr)

	// Start the server
	server.ListenAndServe()
}

type Middleware func(next http.Handler) http.Handler

func MiddlewareChain(middleware ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		for i := len(middleware) - 1; i >= 0; i-- {
			next = middleware[i](next)
		}

		return next
	}
}
