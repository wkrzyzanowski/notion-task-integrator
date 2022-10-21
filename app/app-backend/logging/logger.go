package logging

import (
	"log"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger = setupLogger()

func setupLogger() *zap.Logger {
	if os.Getenv("ENV") == "DEVELOPMENT" {
		return setupLoggerDevelopment()
	}
	return setupLoggerProduction()
}

func setupLoggerDevelopment() *zap.Logger {
	loggerConfig := zap.NewDevelopmentConfig()
	loggerConfig.EncoderConfig.TimeKey = "timestamp"
	loggerConfig.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)

	logging, err := loggerConfig.Build()
	if err != nil {
		log.Fatal(err)
	}

	logging.Info("Logger setup: DEVELOPMENT!")

	return logging
}

func setupLoggerProduction() *zap.Logger {
	loggerConfig := zap.NewProductionConfig()
	loggerConfig.EncoderConfig.TimeKey = "timestamp"
	loggerConfig.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)

	logging, err := loggerConfig.Build()
	if err != nil {
		log.Fatal(err)
	}

	logging.Info("Logger setup: PRODUCTION!")

	return logging
}
