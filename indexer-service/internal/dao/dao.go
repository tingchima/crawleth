// Package dao provides
package dao

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"portto_interview/indexer-service/pkg/storage"
)

// Dao .
type Dao struct {
	db *storage.Database

	cache storage.Cache
}

// NewDao .
func NewDao(db *storage.Database, cache storage.Cache) *Dao {
	return &Dao{
		db:    db,
		cache: cache,
	}
}

// DB .
func (d *Dao) DB() *storage.Database {
	return d.db
}

// Cache .
func (d *Dao) Cache() storage.Cache {
	return d.cache
}

// Begin .
func (d *Dao) Begin(ctx context.Context) *Dao {
	conn := d.DB().Engine().Begin(&sql.TxOptions{})
	database := &storage.Database{}
	database.WithBegin(conn)
	return &Dao{db: database}
}

// Commit .
func (d *Dao) Commit() error {
	err := d.DB().Engine().Commit().Error
	if err != nil {
		msg := fmt.Sprintf("transaction commit failed reason %v.", err)
		return errors.New(msg)
	}
	return nil
}

// Rollback .
func (d *Dao) Rollback() error {
	return d.DB().Engine().Rollback().Error
}
