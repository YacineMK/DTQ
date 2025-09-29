package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/YacineMK/DTQ/services/trip-service/internal/domain"
	"github.com/YacineMK/DTQ/services/trip-service/internal/infrastructure/repository"
	"github.com/YacineMK/DTQ/services/trip-service/utils"
	trippb "github.com/YacineMK/DTQ/shared/proto/trip"
)

type TripService struct {
	trippb.UnimplementedTripServiceServer
	repo repository.MongoRepository
}

func NewTripService(r repository.MongoRepository) *TripService {
	return &TripService{
		repo: r,
	}
}

func (s *TripService) PreviewTrip(ctx context.Context, req *trippb.PreviewTripRequest) (*trippb.PreviewTripResponse, error) {
	if req == nil {
		return nil, errors.New("request cannot be nil")
	}

	packageSlug := req.GetPackageSlug()
	if packageSlug == "" {
		return nil, errors.New("packageSlug is required")
	}

	start := req.GetStartLocation()
	end := req.GetEndLocation()
	if start == nil || end == nil {
		return nil, errors.New("both start and end locations are required")
	}

	routeString := fmt.Sprintf("%f,%f|%f,%f",
		start.Latitude, start.Longitude,
		end.Latitude, end.Longitude,
	)

	result, err := utils.GetRoute(routeString)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch route: %w", err)
	}

	if len(result.Routes) == 0 {
		return nil, errors.New("no routes found from OSRM")
	}

	firstRoute := result.Routes[0]
	if firstRoute.Distance <= 0 || firstRoute.Duration <= 0 {
		return nil, errors.New("invalid route data received")
	}

	rideFare := &trippb.RideFare{
		PackageSlug: packageSlug,
		TotalPrice:  firstRoute.Duration * 40, // fare calculation example
		Route: &trippb.Route{
			Distance: firstRoute.Distance,
			Duration: firstRoute.Duration,
			Coordinates: []*trippb.Coordinate{
				start,
				end,
			},
		},
	}

	return &trippb.PreviewTripResponse{
		RideFares: rideFare,
	}, nil
}

func (s *TripService) CreateTrip(ctx context.Context, req *trippb.CreateTripRequest) (*trippb.Trip, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}

	userID := req.GetUserID()
	if userID == "" {
		return nil, fmt.Errorf("userID is required")
	}

	rideFareID := req.GetRideFareID()
	if rideFareID == "" {
		return nil, fmt.Errorf("rideFareID is required")
	}

	rideFare, err := s.repo.GetRideFareByID(ctx, rideFareID)
	if err != nil {
		return nil, fmt.Errorf("failed to get ride fare %q: %w", rideFareID, err)
	}

	const statusCreated = domain.TripCreated

	tripModel := &domain.TripModel{
		UserID:   userID,
		RideFare: rideFare,
		Status:   statusCreated,
	}

	trip, err := s.repo.CreateTrip(ctx, tripModel)
	if err != nil {
		return nil, fmt.Errorf("failed to create trip: %w", err)
	}

	return &trippb.Trip{
		Id:     trip.ID.Hex(),
		UserID: userID,
		Status: string(trip.Status),
	}, nil
}
