package logger

import (
	"os"

	"go.uber.org/zap"
)

var log *zap.Logger

func Init() {
	var err error
	env := os.Getenv("APP_ENV")

	if env == "dev" {
		log, err = zap.NewDevelopment()
	} else {
		log, err = zap.NewProduction()
	}
	if err != nil {
		panic(err)
	}
}

func Sync() {
	_ = log.Sync()
}

func Info(msg string, fields ...zap.Field) {
	log.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	log.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	log.Error(msg, fields...)
}

func Debug(msg string, fields ...zap.Field) {
	log.Debug(msg, fields...)
}
