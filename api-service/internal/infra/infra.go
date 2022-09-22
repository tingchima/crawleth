// Package infra provides
package infra

import (
	"log"

	"portto_interview/api-service/global"
	"portto_interview/api-service/pkg/config"
	"portto_interview/api-service/pkg/logger"
	"portto_interview/api-service/pkg/storage"

	"github.com/rs/zerolog"
)

// SetupInfra .
func SetupInfra() error {
	err := setupConfig()
	if err != nil {
		log.Fatalf("setup config err: %v\n", err)
	}
	err = setupDatabase()
	if err != nil {
		log.Fatalf("setup database err: %v\n", err)
	}
	err = setupLogger()
	if err != nil {
		log.Fatalf("setup logger err: %v\n", err)
	}
	return nil
}

// setupConfig .
func setupConfig() error {
	cfg, err := config.NewConfig("app", "configs/")
	if err != nil {
		return err
	}
	err = cfg.ReadConfig("App", &global.AppConfig)
	if err != nil {
		return err
	}
	err = cfg.ReadConfig("HttpServer", &global.HTTPServerConfig)
	if err != nil {
		return err
	}
	err = cfg.ReadConfig("Database", &global.DatabaseConfig)
	if err != nil {
		return err
	}
	// Add others config at here
	return nil
}

// setupDatabase .
func setupDatabase() error {
	var err error
	global.Database, err = storage.NewDatabase(global.DatabaseConfig)
	if err != nil {
		return err
	}
	return nil
}

// setupLogger .
func setupLogger() error {
	zLog := zerolog.Logger{}
	global.Logger = logger.NewLogger(&zLog, "", log.LstdFlags).WithCaller(2)
	return nil
}
