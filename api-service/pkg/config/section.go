// Package config provides
package config

import "time"

// AppConfig .
type AppConfig struct {
	CrawlerSize    int
	MaxCrawlerSize int

	DefaultPageSize int
	MaxPageSize     int
}

// EthereumConfig .
type EthereumConfig struct {
	URL string
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

// HTTPServerConfig .
type HTTPServerConfig struct {
	RunMode      string
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}
