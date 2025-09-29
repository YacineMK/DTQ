package repository

import (
	"context"

	"github.com/YacineMK/DTQ/services/trip-service/internal/domain"
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

func (d *MongoRepository) CreateTrip(ctx context.Context, trip *domain.TripModel) (*domain.TripModel, error) {
	result, err := d.db.Collection(db.TripsCollection).InsertOne(ctx, trip)
	if err != nil {
		return nil, err
	}

	trip.ID = result.InsertedID.(primitive.ObjectID)
	return trip, nil
}

func (d *MongoRepository) GetTripByID(ctx context.Context, id string) (*domain.TripModel, error) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	result := d.db.Collection(db.TripsCollection).FindOne(ctx, bson.M{"_id": _id})
	if result.Err() != nil {
		return nil, result.Err()
	}

	var trip domain.TripModel
	if err := result.Decode(&trip); err != nil {
		return nil, err
	}

	return &trip, nil
}

func (d *MongoRepository) UpdateTrip(ctx context.Context, tripID string, status *domain.TripStatus) error {
	_id, err := primitive.ObjectIDFromHex(tripID)
	if err != nil {
		return err
	}

	update := bson.M{"$set": bson.M{"status": status}}

	_, err = d.db.Collection(db.TripsCollection).UpdateOne(ctx, bson.M{"_id": _id}, update)
	if err != nil {
		return err
	}

	return nil
}

func (d *MongoRepository) SaveRideFare(ctx context.Context, fare *domain.RideFareModel) error {
	result, err := d.db.Collection(db.RideFaresCollection).InsertOne(ctx, fare)
	if err != nil {
		return err
	}

	fare.ID = result.InsertedID.(primitive.ObjectID).Hex()
	return nil
}

func (d *MongoRepository) GetRideFareByID(ctx context.Context, id string) (*domain.RideFareModel, error) {
	result := d.db.Collection(db.RideFaresCollection).FindOne(ctx, bson.M{"_id": id})
	if result.Err() != nil {
		return nil, result.Err()
	}

	var fare domain.RideFareModel
	if err := result.Decode(&fare); err != nil {
		return nil, err
	}

	return &fare, nil
}