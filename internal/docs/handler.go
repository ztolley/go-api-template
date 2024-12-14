package docs

import (
	"embed"
	"net/http"
	"strings"
	"time"

	httpSwagger "github.com/swaggo/http-swagger"
)

//go:embed openapi.yaml openapi.json
var content embed.FS

// RegisterRoutes registers the routes for serving the OpenAPI documentation (embedded)
// in the binary, and Swagger UI using the swagger http package
func RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("/openapi.yaml", serveFile("openapi.yaml"))
	router.HandleFunc("/openapi.json", serveFile("openapi.json"))
	router.Handle("/swagger/", httpSwagger.Handler(
		httpSwagger.URL("/openapi.json"),
	))
}

func serveFile(filename string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := content.ReadFile(filename)
		if err != nil {
			http.Error(w, "File not found", http.StatusNotFound)
			return
		}
		http.ServeContent(w, r, filename, time.Time{}, strings.NewReader(string(data)))
	}
}
