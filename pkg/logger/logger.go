package logger

import (
	"go.uber.org/zap"

	"collab/pkg/config"
)

// InitSugarLogger returns a Zap SugaredLogger.
func InitSugarLogger() *zap.SugaredLogger {
	var logger *zap.Logger
	var err error
	if config.GetAppEnv() == "dev" {
		logger, err = zap.NewDevelopment()
	} else {
		logger, err = zap.NewProduction()
	}
	if err != nil {
		panic("Cannot initialize zap logger")
	}

	return logger.Sugar()
}
