package main

import (
	"github.com/Jamshidbek1608/temp-service/config"
	"github.com/Jamshidbek1608/temp-service/pkg/db"
	"github.com/Jamshidbek1608/temp-service/pkg/logger"
	"github.com/Jamshidbek1608/temp-service/service"
)

func main() {
	cfg := config.Load()

	log := logger.New(cfg.LogLevel, "example-service")
	defer logger.Cleanup(log)

	log.Info("main: sqlxConfig",
		logger.String("host", cfg.PostgresHost),
		logger.Int("port", cfg.PostgresPort),
		logger.String("database", cfg.PostgresDatabase))

	connDB, err := db.ConnectToDB(cfg)
	if err != nil {
		log.Fatal("sqlx connection to postgres error", logger.Error(err))
	}

	userService := service.NewUserService

}
