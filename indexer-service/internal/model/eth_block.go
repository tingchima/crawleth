// Package model provides
package model

import (
	"portto_interview/indexer-service/pkg/storage"
	"time"

	"github.com/shopspring/decimal"
)

// EthBlock .
type EthBlock struct {
	Number        decimal.Decimal  `json:"number"`
	HashHex       string           `json:"hash_hex"`
	Time          int64            `json:"time"`
	ParentHashHex string           `json:"parent_hash_hex"`
	IsStable      int8             `json:"is_stable"`
	Transactions  []EthTransaction `json:"transactions,omitempty" gorm:"foreignKey:block_hash_hex;references:hash_hex"`
}

const (
	// BlockIsNotStable .
	BlockIsNotStable int8 = iota
	// BlockIsStable .
	BlockIsStable
)

// TableName .
func (b EthBlock) TableName() string {
	return "eth_blocks"
}

// EthBlockUpdates .
type EthBlockUpdates struct {
	UpdatedOn time.Time
	IsStable  *int8
}

// EthBlockCondition .
type EthBlockCondition struct {
	BlockNumberLte int64
}

// Create .
func (b EthBlock) Create(db *storage.Database) error {
	return db.Engine().Create(&b).Error
}

// Update .
func (b EthBlock) Update(db *storage.Database, condition *EthBlockCondition, updates *EthBlockUpdates) error {
	tx := db.Engine().Model(&EthBlock{})

	if condition != nil {
		tx = tx.Where("is_stable = ?", b.IsStable)

		if condition.BlockNumberLte != 0 {
			tx = tx.Where("number <= ?", condition.BlockNumberLte)
		}
	}

	updateCols := make(map[string]interface{})
	if updates != nil {
		if updates.IsStable != nil {
			updateCols["is_stable"] = updates.IsStable
		}
		updateCols["updated_on"] = time.Now().UTC()
	}

	err := tx.Updates(updateCols).Error
	if err != nil {
		return err
	}
	return nil
}
