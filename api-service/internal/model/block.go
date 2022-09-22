// Package model provides
package model

import (
	"portto_interview/api-service/pkg/app"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

// Block .
type Block struct {
	Number        decimal.Decimal `json:"number"`
	HashHex       string          `json:"hash_hex"`
	Time          int64           `json:"time"`
	ParentHashHex string          `json:"parent_hash_hex"`
	IsStable      int8            `json:"is_stable"`
	Transactions  []Transaction   `json:"transactions" gorm:"foreignKey:block_hash_hex;references:hash_hex"`
}

// TableName .
func (b Block) TableName() string {
	return "eth_blocks"
}

// BlockOptions block where options
type BlockOptions struct {
	LimitBlockNumber    uint
	Sorting             *app.Sorting // ex. order by "number desc"
	PreloadTransactions bool
}

// BlockSwagger .
type BlockSwagger struct {
	List []*Block
}

// List .
func (b Block) List(db *gorm.DB, c BlockOptions) ([]Block, error) {
	var blocks []Block
	if c.LimitBlockNumber > 0 {
		db = db.Limit(int(c.LimitBlockNumber))
	}
	if c.Sorting != nil {
		db = db.Order(c.Sorting.SortBy())
	}
	if c.PreloadTransactions {
		db = db.Preload("Transactions")
	}
	err := db.Find(&blocks).Error
	if err != nil {
		return nil, err
	}
	return blocks, nil
}

// Get .
func (b Block) Get(db *gorm.DB, c BlockOptions) (block Block, err error) {
	if b.Number != decimal.NewFromInt(0) {
		db = db.Where("number = ?", b.Number.IntPart())
	}
	if c.PreloadTransactions {
		db = db.Preload("Transactions")
	}
	err = db.Find(&block).Error
	if err != nil {
		return
	}
	return
}
