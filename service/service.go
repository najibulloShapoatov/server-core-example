package service

import (
	"github.com/najibulloShapoatov/server-core-example/internal/auth"
	"github.com/najibulloShapoatov/server-core-example/internal/config"
	"github.com/najibulloShapoatov/server-core-example/internal/db"
	"github.com/najibulloShapoatov/server-core/cache/redis"
	"github.com/najibulloShapoatov/server-core/monitoring/log"
	"github.com/najibulloShapoatov/server-core/server"
	"gorm.io/gorm"
)

const (
	moduleID      = "api"
	moduleVersion = "v1"
)

type Config struct {
	Redis    redis.Config          `config:"."`
	DBConfig config.DatabaseConfig `config:"."`
}

type Service struct {
	config *Config
	redis  *redis.Cache
	database     *gorm.DB
}

var svc *Service

// Create and initialize a new API service
func New(c *Config) (*Service, error) {
	if svc != nil {
		return svc, nil
	}

	svc = &Service{
		config: c,
		redis:  redis.New(&c.Redis),
	}

	server.UseMiddleware(auth.Middleware)

	db, err := db.NewConnection(&c.DBConfig)
	if err != nil {
		log.Error(err)
		panic(err)
	}
	svc.database = db
	

	return svc, nil
}

// Implements server-core Service interface. Called before webserver starts
func (s *Service) Setup() error {
	return nil
}

// Implements server-core Service interface. Called when webserver starts
func (s *Service) Start() error {
	return nil
}

// Implements server-core Service interface. Called when webserver stops
func (s *Service) Stop() error {
	return nil
}

// Implements server-core Module interface
func (s *Service) ID() string {
	return moduleID
}

// Implements server-core Module interface
func (s *Service) Version() string {
	return moduleVersion
}
