package main

import (
	"github.com/YacineMK/DTQ/services/trip-service/internal/infrastructure/repository"
	"github.com/YacineMK/DTQ/shared/db"
)

func main() {
	// db 
	cfg := db.NewMongoDefaultConfig()

	client , err := db.NewMongoClient(cfg)
	if err != nil {
		panic(err)
	}

	database := db.GetDatabase(client,cfg)
	repository.NewMongoRepository(database)
}