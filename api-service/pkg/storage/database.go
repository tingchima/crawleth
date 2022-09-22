// Package storage provides
package storage

import (
	"fmt"

	"portto_interview/api-service/pkg/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Database .
type Database struct {
	engine *gorm.DB
}

// Engine .
func (d *Database) Engine() *gorm.DB {
	return d.engine
}

// NewDatabase .
func NewDatabase(cfg *config.DatabaseConfig) (*Database, error) {
	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
	)
	if cfg.SSLEnable {
		dsn += " sslmode=require"
	} else {
		dsn += " sslmode=disable"
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	fmt.Println("database ping success")
	return &Database{engine: db}, nil
}
