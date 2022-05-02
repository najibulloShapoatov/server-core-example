package config

import "time"

type DatabaseConfig struct {
	Addr                   string        `config:"platform.db.addr" default:"localhost"`
	Name                   string        `config:"platform.db.name" default:"testdb"`
	Username               string        `config:"platform.db.username" default:"postgres"`
	Password               string        `config:"platform.db.password" default:"postgres"`
	ConnMaxIdleTime        time.Duration `config:"platform.db.connMaxIdleTime" default:"30m"`
	ConnectionMaxLifetime  time.Duration `config:"platform.db.connectionMaxLifetime" default:"3d"`
	Port                   uint16        `config:"platform.db.port" default:"5432"`
	MaxIDLEConnectionCount int           `config:"platform.db.maxIDLEConnectionCount" default:"15"`
	MaxOpenConnectionCount int           `config:"platform.db.maxOpenConnectionCount" default:"8"`
}
