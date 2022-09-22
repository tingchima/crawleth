// Package infra provides
package infra

import (
	"context"
	"log"

	"portto_interview/indexer-service/global"
	"portto_interview/indexer-service/pkg/config"
	"portto_interview/indexer-service/pkg/logger"
	"portto_interview/indexer-service/pkg/storage"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/panjf2000/ants/v2"
	"github.com/rs/zerolog"
)

// SetupInfra .
func SetupInfra() error {
	err := setupConfig()
	if err != nil {
		log.Fatalf("setup config err: %v\n", err)
	}
	err = setupEthclient()
	if err != nil {
		log.Fatalf("setup ethclient err: %v\n", err)
	}
	err = setupCrawlerPool()
	if err != nil {
		log.Fatalf("setup crawler pool err: %v\n", err)
	}
	err = setupDatabase()
	if err != nil {
		log.Fatalf("setup database err: %v\n", err)
	}
	err = setupRedis()
	if err != nil {
		log.Fatalf("setup redis err: %v\n", err)
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
	err = cfg.ReadConfig("AntsConfig", &global.AntsConfig)
	if err != nil {
		return err
	}
	err = cfg.ReadConfig("Ethereum", &global.EthereumConfig)
	if err != nil {
		return err
	}
	err = cfg.ReadConfig("Database", &global.DatabaseConfig)
	if err != nil {
		return err
	}
	err = cfg.ReadConfig("Redis", &global.RedisConfig)
	if err != nil {
		return err
	}
	// Add others config at here
	return nil
}

// setupEthclient .
func setupEthclient() error {
	var err error
	global.EthereumRPC, err = ethclient.Dial(global.EthereumConfig.RPC)
	if err != nil {
		return err
	}
	chainID, err := global.EthereumRPC.NetworkID(context.Background())
	if err != nil {
		return err
	}
	global.Ethsigner = types.NewEIP155Signer(chainID)
	return nil
}

// setupCrawlerPool .
func setupCrawlerPool() error {
	var err error
	global.EthCrawlerPool, err = ants.NewPool(global.AntsConfig.PoolSize)
	if err != nil {
		return err
	}
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

// setupRedis .
func setupRedis() error {
	var err error
	global.Cache, err = storage.NewRedis(global.RedisConfig)
	if err != nil {
		return err
	}
	return err
}

// setupLogger .
func setupLogger() error {
	zLog := zerolog.Logger{}
	global.Logger = logger.NewLogger(&zLog, "", log.LstdFlags).WithCaller(2)
	return nil
}
