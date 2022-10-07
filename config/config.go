package config

import (
	"os"

	"github.com/spf13/cast"
)

type Config struct {
	Enviroment        string
	PostgresHost      string
	PostgresPost      int
	PostgresDatabase  string
	PostgresUser      string
	PostgresPassword  string
	LogLevel          string
	RPCPort           string
	ReviewServiceHost string
	ReviewServicePort int
}

func Load() Config {
	c := Config{}

	c.Enviroment = cast.ToString(getOReturnDefault("ENVIRONMENT", "develop"))

	c.PostgresHost = cast.ToString(getOReturnDefault("POSTGRES_HOST", "localhost"))
	c.PostgresPost = cast.ToInt(getOReturnDefault("POSTGRES_PORT", 5432))
	c.PostgresDatabase = cast.ToString(getOReturnDefault("POSTGRES_DATABASE", "userdb"))
	c.PostgresUser = cast.ToString(getOReturnDefault("POSTGRES_USER", "postgres"))
	c.PostgresPassword = cast.ToString(getOReturnDefault("POSTGRES_PASSWORD", 123))

	c.LogLevel = cast.ToString(getOReturnDefault("LOG_LEVEL", "debug"))

	c.RPCPort = cast.ToString(getOReturnDefault("RPC_PORT", ":8080"))

	return c
}

func getOReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}
	return defaultValue
}
