package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	driverpb "github.com/YacineMK/DTQ/shared/proto/driver"

)

type DriverStatus string

const (
	DriverAvailable   DriverStatus = "available"
	DriverUnavailable DriverStatus = "unavailable"
)

type DriverModel struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Name   string             `bson:"name"`
	Email  string             `bson:"email"`
	Status DriverStatus       `bson:"status"`
}

// repository
type DriverRepository interface {
	GetAvailableDriver(ctx context.Context) (*DriverModel, error)
	UpdateDriverStatus(ctx context.Context, driverID string, status DriverStatus) error
	CreateDriver(ctx context.Context, driver *DriverModel) (*DriverModel, error)
}

// service
type DriverService interface {
	CreateDriver(ctx context.Context, req *driverpb.RegisterDriverRequest) (*driverpb.RegisterDriverResponse, error)
}
