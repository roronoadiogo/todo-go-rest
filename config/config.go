package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

const loggerLevel = zap.InfoLevel

var logger *zap.SugaredLogger

func init() {
	config := zap.NewProductionConfig()
	config.Level.SetLevel(loggerLevel)

	var err error
	logger, err := config.Build()
	if err != nil {
		panic(err)
	}

	logger.Info("Getting the Enviroment definitions")

	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "dev"
	}

	err = godotenv.Load(fmt.Sprintf("config/.env.%s", env))

	if err != nil {
		logger.Error("Error in env params definitios, application should be stopped")
		panic(err)
	}

	if err = loadEnv(); err != nil {
		logger.Sugar().Errorw("Error loading environment variables", "error", err)
		panic(err)
	}

}

func loadEnv() error {

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
