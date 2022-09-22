// Package global provides
package global

import "portto_interview/indexer-service/pkg/storage"

var (
	// Database .
	Database *storage.Database
	// Cache .
	Cache storage.Cache
)
