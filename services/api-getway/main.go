package main

import (
	"github.com/YacineMK/DTQ/shared/env"
	"github.com/YacineMK/DTQ/shared/logger"
)

var (
	http_addr = env.GetEnv("HTTP_ADRR", ":8080")
)

func main() {
	logger.Init()
	logger.Info("Starting API Gateway")
}
