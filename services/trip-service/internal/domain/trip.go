package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TripStatus string

const (
	TripCreated        TripStatus = "created"
	TripDriverAssigned TripStatus = "driver_assigned"
	TripCompleted      TripStatus = "completed"
	TripCancelled      TripStatus = "cancelled"
)

type TripModel struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	UserID   string             `bson:"userID"`
	Status   TripStatus         `bson:"status"`
	RideFare *RideFareModel     `bson:"rideFare"`
}

// repository
type TripRepository interface {
	CreateTrip(ctx context.Context, trip *TripModel) (*TripModel, error)
	GetTripByID(ctx context.Context, id string) (*TripModel, error)
	UpdateTrip(ctx context.Context, tripID string, status TripStatus) error
	SaveRideFare(ctx context.Context, tripID string, fare *RideFareModel) error
}

// service
type TripService interface {
	CreateTrip(ctx context.Context, fare *RideFareModel) (*TripModel, error)
}