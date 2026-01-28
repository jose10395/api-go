package database

import (
	config "api-go/src"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresDB(cfg config.Config) (*gorm.DB, error) {
	return gorm.Open(
		postgres.Open(cfg.PostgresDSN),
		&gorm.Config{},
	)
}
