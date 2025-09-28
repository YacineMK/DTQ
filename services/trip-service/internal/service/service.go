package service

import (
    "context"

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
    packageSlug := req.GetPackageSlug()
    start := req.GetStartLocation()
    end := req.GetEndLocation()

    routeString := start.String() + "|" + end.String()
    result, err := utils.GetRoute(routeString)
    if err != nil {
        return nil, err
    }

    rideFare := &trippb.RideFare{
        PackageSlug: packageSlug,
        TotalPrice:  result.Routes[0].Duration * 40,
        Route: &trippb.Route{
            Distance:   result.Routes[0].Distance,
            Duration:   result.Routes[0].Duration,
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
