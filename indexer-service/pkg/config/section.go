// Package config provides
package config

import "time"

// AppConfig .
type AppConfig struct {
	EthSyncingDuration   time.Duration
	EthUpdateBlockStatus time.Duration
	EthStableBlockRule   int64
	EthCrawlerSize       int64
	MaxSyncHeaderSize    int64

	DefaultPageSize int
	MaxPageSize     int
}

// AntsConfig .
type AntsConfig struct {
	PoolSize int
}

// EthereumConfig .
type EthereumConfig struct {
	RPC                string
	SubscriberDuration time.Duration
}

// DatabaseConfig .
type DatabaseConfig struct {
	Username     string
	Password     string
	Host         string
	Port         int
	DBName       string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
	SSLEnable    bool
}

// RedisConfig .
type RedisConfig struct {
	Addresses []string
	Password  string
	DB        int
}

// HTTPServerConfig .
type HTTPServerConfig struct {
	RunMode      string
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

// NatsConfig .
type NatsConfig struct {
	URL string
}
