package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/YacineMK/DTQ/services/driver-service/internal/domain"
	"github.com/YacineMK/DTQ/services/driver-service/internal/infrastructure/repository"
	driverpb "github.com/YacineMK/DTQ/shared/proto/driver"
)

type DriverService struct {
	driverpb.UnimplementedDriverServiceServer
	repo repository.MongoRepository
}

func NewDriverService(r repository.MongoRepository) *DriverService {
	return &DriverService{
		repo: r,
	}
}
func (s *DriverService) CreateDriver(ctx context.Context, req *driverpb.RegisterDriverRequest) (*driverpb.RegisterDriverResponse, error) {
	if req == nil {
		return nil, errors.New("request is nil")
	}
	userMail := req.GetEmail();
	if userMail == ""  {
		return nil, errors.New("email is empty")
	}

	userName :=  req.GetName();
	if userName == ""  {
		return nil, errors.New("name is empty")
	}

	const StatusAvilable = domain.DriverAvailable

	driverModel := &domain.DriverModel{
		Name: userName,
		Email: userMail,
		Status: StatusAvilable,
	}

	_ , err := s.repo.CreateDriver(ctx,driverModel)
	if err != nil {
		return nil , fmt.Errorf("failed to create driver: %w", err)
	}
	return &driverpb.RegisterDriverResponse{
		Msg: "driver created successfully",
	}, nil
}
