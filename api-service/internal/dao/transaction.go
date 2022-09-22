// Package dao provides
package dao

import (
	"context"

	"portto_interview/api-service/internal/model"
)

// GetTransactionByHash .
func (d *Dao) GetTransactionByHash(ctx context.Context, hashHex string) (model.Transaction, error) {
	transaction := model.Transaction{HashHex: hashHex}
	condition := model.TransactionOptions{
		PreloadLogs: true,
	}
	return transaction.Get(d.db.Engine(), condition)
}
