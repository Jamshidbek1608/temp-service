package main

import (
	"temp-service/pkg/logger"

	"temp-service/config"
)

func main() {
	cfg := config.Load()

	log := logger.New(cfg.LogLevel, "example-service")
	defer logger.Cleanup(log)

	log.Info("main: sqlxConfig",
		logger.String("host", cfg.PostgresHost),
		logger.Int("port", cfg.PostgresPort),
		logger.String("database", cfg.PostgresDatabase))

}
