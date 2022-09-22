// Package model provides
package model

// TransactionLog .
type TransactionLog struct {
	TxHashHex string `json:"tx_hash_hex"`
	LogIndex  uint   `json:"log_index"`
	LogData   []byte `json:"log_data"`
}

// TableName .
func (t TransactionLog) TableName() string {
	return "eth_transaction_logs"
}
