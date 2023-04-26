package applogger

import (
	"github.com/mfereji-io/devkit-go/example/kaska-svc-api/internal/config"
	"go.uber.org/zap"
)

type (
	AppLogger struct {
		Logger *zap.Logger
	}
)

func InitAppLogger(c *config.AppConfig) *zap.Logger {

	if c.AppEnv == "prod" {

		return InitAppProdLogger()

	} else{
		return InitAppDevLogger()

	} 
}

func InitAppProdLogger() *zap.Logger {

	logger, err := zap.NewProduction()

	if err != nil {
		panic("could not init app logger:prod")
	}

	return logger
}

func InitAppDevLogger() *zap.Logger {

	logger, err := zap.NewDevelopment()

	if err != nil {
		panic("could not init app logger:dev")
	}

	return logger
}
