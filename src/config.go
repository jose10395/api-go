package config

import "os"

type Config struct {
	PostgresDSN string
	RedisAddr   string
	RedisPass   string
}

func LoadConfigFromEnv() Config {
	return Config{
		PostgresDSN: os.Getenv("POSTGRES_DSN"),
		RedisAddr:   os.Getenv("REDIS_ADDR"),
		RedisPass:   os.Getenv("REDIS_PASS"),
	}
}
