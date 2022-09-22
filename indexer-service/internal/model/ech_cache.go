// Package model provides
package model

import (
	"math/big"
	"sync"
)

// EthCache .
type EthCache struct {
	sync.RWMutex
	blocks          []EthBlock
	txs             []EthTransaction
	txLogs          []EthTransactionLog
	errBlockNumbers []*big.Int
	c               chan int
}

// SetBlock .
func (c *EthCache) SetBlock(b EthBlock) {
	c.Lock()
	c.blocks = append(c.blocks, b)
	c.Unlock()
	con := sync.NewCond(&c.RWMutex)
	con.Broadcast()
}

// SetTransactions .
func (c *EthCache) SetTransactions(txs []EthTransaction) {
	c.Lock()
	c.txs = append(c.txs, txs...)
	c.Unlock()
}

// SetTxLogs .
func (c *EthCache) SetTxLogs(logs []EthTransactionLog) {
	c.Lock()
	c.txLogs = append(c.txLogs, logs...)
	c.Unlock()
}

// GetBlocks .
func (c *EthCache) GetBlocks() []EthBlock {
	return c.blocks
}

// GetTransactions .
func (c *EthCache) GetTransactions() []EthTransaction {
	return c.txs
}

// GetTransactionLogs .
func (c *EthCache) GetTransactionLogs() []EthTransactionLog {
	return c.txLogs
}
