package db

import (
	"context"
	"time"

	"github.com/YacineMK/DTQ/shared/env"
	"github.com/YacineMK/DTQ/shared/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.uber.org/zap"
)

var (
	uri = env.GetEnv("MONGO_DB", "mongodb://localhost:27017")
)

const (
	TripsCollection = "trips"
	RideFaresCollection = "ride_fares"
)

type MongoConfig struct {
	URI      string
	Database string
}

func NewMongoDefaultConfig() *MongoConfig {
	return &MongoConfig{
		URI:      uri,
		Database: "DTQ",
	}
}

func NewMongoClient(cfg *MongoConfig) (*mongo.Client, error) {
	logger.Init()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.URI))
	if err != nil {
		logger.Error("Failed to create MongoDB client", zap.Error(err))
		return nil, err
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		logger.Error("Failed to ping MongoDB", zap.Error(err))
		return nil, err
	}

	logger.Info("Successfully connected to MongoDB", zap.String("uri", cfg.URI))
	return client, nil
}

func GetDatabase(client *mongo.Client, cfg *MongoConfig) *mongo.Database {
	return client.Database(cfg.Database)
}
