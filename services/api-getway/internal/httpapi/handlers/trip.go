package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/YacineMK/DTQ/services/api-getway/internal/grpcclients"
	trippb "github.com/YacineMK/DTQ/shared/proto/trip"
)

func HandlePreviewTrip(w http.ResponseWriter, r *http.Request, tripClient *grpcclients.TripService) {
	var req trippb.PreviewTripRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	resp, err := tripClient.Client.PreviewTrip(ctx, &req)
	if err != nil {
		http.Error(w, "grpc error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(resp)
}

func HandleCreateTrip(w http.ResponseWriter, r *http.Request, tripClient *grpcclients.TripService) {
	var req trippb.CreateTripRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	resp, err := tripClient.Client.CreateTrip(ctx, &req)
	if err != nil {
		http.Error(w, "grpc error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(resp)
}
