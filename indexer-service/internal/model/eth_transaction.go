// Package model provides
package model

import (
	"portto_interview/indexer-service/pkg/storage"
)

// EthTransaction .
type EthTransaction struct {
	BlockHashHex string              `json:"block_hash_hex"`
	HashHex      string              `json:"hash_hex"`
	FromHex      string              `json:"from_hex"`
	ToHex        string              `json:"to_hex"`
	Nonce        uint64              `json:"nonce"`   // the number of transactions made by the sender prior to this one
	TxData       []byte              `json:"tx_data"` // the data send along with the transaction
	Value        string              `json:"value"`   // value transferred in Wei
	Logs         []EthTransactionLog `json:"logs,omitempty" gorm:"foreignKey:tx_hash_hex;references:hash_hex"`
}

// TableName .
func (t EthTransaction) TableName() string {
	return "eth_transactions"
}

// EthTransactions .
type EthTransactions []EthTransaction

// BatchCreate .
func (t EthTransactions) BatchCreate(db *storage.Database) error {
	return db.Engine().CreateInBatches(t, len(t)).Error
}
