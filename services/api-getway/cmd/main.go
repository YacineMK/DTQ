package main

import (
	"log"
	"net/http"

	"github.com/YacineMK/DTQ/services/gateway/internal/grpcclients"
	"github.com/YacineMK/DTQ/services/gateway/internal/httpapi"
	"github.com/YacineMK/DTQ/shared/env"
	"github.com/YacineMK/DTQ/shared/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	httpAddr    = env.GetEnv("HTTP_ADDR", ":8080")
	tripSvcAddr = env.GetEnv("TRIP_SERVICE_ADDR", ":50051")
)

func main() {
	logger.Init()
	logger.Info("Starting API Gateway")

	tripClient, err := grpcclients.NewTripServiceClient(tripSvcAddr)
	if err != nil {
		log.Fatalf("failed to connect to trip-service: %v", err)
	}
	defer tripClient.Conn.Close()

	mux := http.NewServeMux()
	httpapi.RegisterRoutes(mux, tripClient)

	logger.Info("HTTP Gateway listening on", zap.String("address", httpAddr))
	if err := http.ListenAndServe(httpAddr, mux); err != nil {
		log.Fatal(err)
	}
}
