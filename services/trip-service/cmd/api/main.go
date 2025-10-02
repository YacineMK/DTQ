package main

import (
	"net"

	"github.com/YacineMK/DTQ/services/trip-service/internal/infrastructure/repository"
	"github.com/YacineMK/DTQ/services/trip-service/internal/service"
	"github.com/YacineMK/DTQ/shared/db"
	"github.com/YacineMK/DTQ/shared/env"
	trippb "github.com/YacineMK/DTQ/shared/proto/trip"
	"google.golang.org/grpc"
)

var (
	tripSrvPort = env.GetEnv("TRIP_SERVICE_PORT", ":50051")
)

func main() {
	// db
	cfg := db.NewMongoDefaultConfig()

	client, err := db.NewMongoClient(cfg)
	if err != nil {
		panic(err)
	}

	database := db.GetDatabase(client, cfg)
	repos := repository.NewMongoRepository(database)

	tripService := service.NewTripService(*repos)

	lis, err := net.Listen("tcp", tripSrvPort)
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()

	trippb.RegisterTripServiceServer(grpcServer, tripService)

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}

}
