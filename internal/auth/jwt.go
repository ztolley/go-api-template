package auth

import (
	"log"
	"net/http"
	"strings"

	"github.com/ztolley/goapi/internal/utils"
)

// WithJWTAuth returns a middleware that enforces JWT authentication, excluding specified routes
func WithJWTAuth(excludedRoutes []string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Bypass authentication for specific paths
			for _, route := range excludedRoutes {
				if strings.HasPrefix(r.URL.Path, route) {
					next.ServeHTTP(w, r)
					return
				}
			}

			tokenString := utils.GetTokenFromRequest(r)

			if tokenString == "" || !validateJWT(tokenString) {
				log.Printf("Unauthorized")
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func validateJWT(_token string) bool {
	return true
}
