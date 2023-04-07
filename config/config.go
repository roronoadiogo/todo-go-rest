package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

const loggerLevel = zap.DebugLevel

var logger *zap.SugaredLogger

func InitializeLogger() *zap.Logger {
	config := zap.NewProductionConfig()
	config.Level.SetLevel(loggerLevel)

	logger, err := config.Build()
	if err != nil {
		panic(err)
	}

	return logger
}

func LoadEnv() error {

	logger.Info("Getting the Enviroment definitions")

	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "dev"
	}

	err := godotenv.Load(fmt.Sprintf("config/.env.%s", env))

	if err != nil {
		logger.Errorw("Error in env params definitions, application should be stopped", "error", err)
		panic(err)
	}

	return nil
}
