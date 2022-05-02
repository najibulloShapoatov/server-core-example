package db

import (
	"fmt"

	"github.com/najibulloShapoatov/server-core-example/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewConnection(config *config.DatabaseConfig) (*gorm.DB, error) {
	connection, err := gorm.Open(
		postgres.Open(fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
			config.Addr,
			config.Username,
			config.Password,
			config.Name,
			config.Port,
		)),
		&gorm.Config{},
	)

	if err != nil {
		return nil, err
	}

	database, err := connection.DB()

	if err != nil {
		return nil, err
	}

	database.SetConnMaxIdleTime(config.ConnMaxIdleTime)
	database.SetConnMaxLifetime(config.ConnectionMaxLifetime)
	database.SetMaxIdleConns(config.MaxIDLEConnectionCount)
	database.SetMaxOpenConns(config.MaxOpenConnectionCount)

	return connection, nil
}
