package config

import (
	"os"
	"strconv"
)

type Config struct {
	PostgresDSN string
	RedisAddr   string
	RedisPass   string
	RedisDB     int
}

func LoadConfigFromEnv() Config {

	redisDB := 0
	if dbStr := os.Getenv("REDIS_DB"); dbStr != "" {
		if db, err := strconv.Atoi(dbStr); err == nil {
			redisDB = db
		}
	}

	return Config{
		PostgresDSN: os.Getenv("POSTGRES_DSN"),
		RedisAddr:   os.Getenv("REDIS_ADDR"),
		RedisPass:   os.Getenv("REDIS_PASS"),
		RedisDB:     redisDB,
	}
}
