// Package model provides
package model

// EthTransactionLog .
type EthTransactionLog struct {
	TxHashHex string `json:"tx_hash_hex"`
	LogIndex  uint   `json:"log_index"`
	LogData   []byte `json:"log_data"`
}

// TableName .
func (t EthTransactionLog) TableName() string {
	return "eth_transaction_logs"
}
