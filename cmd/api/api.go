package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/ztolley/goapi/services/auth"
	"github.com/ztolley/goapi/services/user"
	"github.com/ztolley/goapi/utils"
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

	// Create a user store, that provides a repository to access user data
	userStore := user.NewStore(s.db)

	// Create a new handler service for user related requests, pass it a reference
	// to the user store so it can make database repository calls then register the
	// routes it will handle
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(router)

	// Define the middleware chain for all requests
	middlewareChain := MiddlewareChain(utils.RequestLoggerMiddleware, auth.WithJWTAuth)

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
