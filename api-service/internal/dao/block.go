// Package dao provides
package dao

import (
	"context"

	"portto_interview/api-service/internal/model"
	"portto_interview/api-service/pkg/app"

	"github.com/shopspring/decimal"
)

// GetBlockListByLimit .
func (d *Dao) GetBlockListByLimit(ctx context.Context, limit uint) ([]model.Block, error) {
	block := model.Block{}
	condition := model.BlockOptions{
		LimitBlockNumber:    limit,
		Sorting:             &app.Sorting{Field: "number", Order: app.Desc},
		PreloadTransactions: true,
	}
	return block.List(d.db.Engine(), condition)
}

// GetBlockByNumber .
func (d *Dao) GetBlockByNumber(ctx context.Context, number uint64) (model.Block, error) {
	block := model.Block{Number: decimal.NewFromInt(int64(number))}
	condition := model.BlockOptions{
		PreloadTransactions: true,
	}
	return block.Get(d.db.Engine(), condition)
}
