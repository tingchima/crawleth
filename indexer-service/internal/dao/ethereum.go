// Package dao provides
package dao

import (
	"context"
	"math/big"

	"portto_interview/indexer-service/global"
	"portto_interview/indexer-service/internal/model"
)

// CreateEthBlock .
func (d *Dao) CreateEthBlock(ctx context.Context, block model.EthBlock) error {
	return block.Create(d.DB())
}

// BatchCreateEthBlock .
func (d *Dao) BatchCreateEthBlock(ctx context.Context, blocks []model.EthBlock) error {
	return d.db.Engine().Omit("Transactions").CreateInBatches(blocks, len(blocks)).Error
}

// BatchCreateEthTransaction .
func (d *Dao) BatchCreateEthTransaction(ctx context.Context, transactions []model.EthTransaction) error {
	if transactions == nil || len(transactions) <= 0 {
		return nil
	}
	return d.db.Engine().Omit("Logs").CreateInBatches(transactions, len(transactions)).Error
}

// BatchCreateEthTransactionLog .
func (d *Dao) BatchCreateEthTransactionLog(ctx context.Context, txLogs []model.EthTransactionLog) error {
	if txLogs == nil || len(txLogs) <= 0 {
		return nil
	}
	return d.db.Engine().CreateInBatches(txLogs, len(txLogs)).Error
}

// UpdateEthereumBlockStable .
func (d *Dao) UpdateEthereumBlockStable(ctx context.Context, number *big.Int) error {
	block := model.EthBlock{IsStable: model.BlockIsNotStable}

	condition := &model.EthBlockCondition{
		BlockNumberLte: number.Int64() - global.AppConfig.EthStableBlockRule,
	}
	var isStable int8 = model.BlockIsStable
	updates := &model.EthBlockUpdates{
		IsStable: &isStable,
	}
	err := block.Update(d.DB(), condition, updates)
	if err != nil {
		return err
	}
	return nil
}
