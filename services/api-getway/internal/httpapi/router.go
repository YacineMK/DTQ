package httpapi

import (
	"net/http"

	"github.com/YacineMK/DTQ/services/api-getway/internal/grpcclients"
	"github.com/YacineMK/DTQ/services/api-getway/internal/httpapi/handlers"
)

func RegisterRoutes(mux *http.ServeMux, tripClient *grpcclients.TripService) {
	mux.HandleFunc("/preview-trip", func(w http.ResponseWriter, r *http.Request) {
		handlers.HandlePreviewTrip(w, r, tripClient)
	})
	mux.HandleFunc("/create-trip", func(w http.ResponseWriter, r *http.Request) {
		handlers.HandleCreateTrip(w, r, tripClient)
	})
}
