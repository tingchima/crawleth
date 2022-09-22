// Package global provides
package global

import "portto_interview/api-service/pkg/config"

var (
	// AppConfig .
	AppConfig *config.AppConfig
	// HTTPServerConfig .
	HTTPServerConfig *config.HTTPServerConfig
	// DatabaseConfig .
	DatabaseConfig *config.DatabaseConfig
)
