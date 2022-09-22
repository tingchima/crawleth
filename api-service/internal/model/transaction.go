// Package model provides
package model

import (
	"gorm.io/gorm"
)

// Transaction .
type Transaction struct {
	BlockHashHex string           `json:"block_hash_hex"`
	HashHex      string           `json:"hash_hex"`
	FromHex      string           `json:"from_hex"`
	ToHex        string           `json:"to_hex"`
	Nonce        uint64           `json:"nonce"`   // the number of transactions made by the sender prior to this one
	TxData       []byte           `json:"tx_data"` // the data send along with the transaction
	Value        string           `json:"value"`   // value transferred in Wei
	Logs         []TransactionLog `json:"logs" gorm:"foreignKey:tx_hash_hex;references:hash_hex"`
}

// TableName .
func (t Transaction) TableName() string {
	return "eth_transactions"
}

// TransactionOptions .
type TransactionOptions struct {
	PreloadLogs bool
}

// Get .
func (t Transaction) Get(db *gorm.DB, c TransactionOptions) (tx Transaction, err error) {
	if t.HashHex != "" {
		db = db.Where("hash_hex = ?", t.HashHex)
	}
	if c.PreloadLogs {
		db = db.Preload("Logs")
	}
	err = db.Find(&tx).Error
	if err != nil {
		return
	}
	return
}
