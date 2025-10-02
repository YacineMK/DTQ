package repository

import (
	"context"

	"github.com/YacineMK/DTQ/services/driver-service/internal/domain"
	"github.com/YacineMK/DTQ/shared/db"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepository struct {
	db *mongo.Database
}

func NewMongoRepository(db *mongo.Database) *MongoRepository {
	return &MongoRepository{db: db}
}

func (d *MongoRepository) GetAvailableDriver(ctx context.Context) (*domain.DriverModel, error) {
	results, err := d.db.Collection(db.DriversCollection).Find(ctx, bson.M{"status": domain.DriverAvailable})
	if err != nil {
		return nil, err
	}

	var driver []domain.DriverModel
	if err := results.All(ctx, &driver); err != nil {
		return nil, err
	}
	
	if len(driver) == 0 {
		return nil, nil
	}
	return &driver[0], nil
}

func (d *MongoRepository) UpdateDriverStatus(ctx context.Context, driverID string, status *domain.DriverStatus) error {
	_id, err := primitive.ObjectIDFromHex(driverID)
	if err != nil {
		return err
	}

	update := bson.M{"$set": bson.M{"status": status}}

	_, err = d.db.Collection(db.DriversCollection).UpdateOne(ctx, bson.M{"_id": _id}, update)
	if err != nil {
		return err
	}
	return nil
}

func (d *MongoRepository) CreateDriver(ctx context.Context, driver *domain.DriverModel) (*domain.DriverModel, error) {
	result, err := d.db.Collection(db.DriversCollection).InsertOne(ctx, driver)
	if err != nil {
		return nil, err
	}

	driver.ID = result.InsertedID.(primitive.ObjectID)
	return driver, nil
}