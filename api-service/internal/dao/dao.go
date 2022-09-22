// Package dao provides
package dao

import (
	"portto_interview/api-service/pkg/storage"
)

// Dao .
type Dao struct {
	db *storage.Database
}

// NewDao .
func NewDao(db *storage.Database) *Dao {
	return &Dao{
		db: db,
	}
}
