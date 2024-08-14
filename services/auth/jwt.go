package auth

import (
	"log"
	"net/http"

	"github.com/ztolley/goapi/utils"
)

func WithJWTAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := utils.GetTokenFromRequest(r)

		if tokenString == "" || !validateJWT(tokenString) {
			log.Printf("Unauthorised")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func validateJWT(_token string) bool {
	return true
}
