// Package global provides
package global

import "portto_interview/indexer-service/pkg/config"

var (
	// AppConfig .
	AppConfig *config.AppConfig
	// AntsConfig .
	AntsConfig *config.AntsConfig
	// EthereumConfig .
	EthereumConfig *config.EthereumConfig
	// DatabaseConfig .
	DatabaseConfig *config.DatabaseConfig
	// RedisConfig .
	RedisConfig *config.RedisConfig
)
