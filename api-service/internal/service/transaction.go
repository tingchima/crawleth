// Package service provides
package service

import (
	"context"
)

// GetTransactionReq .
type GetTransactionReq struct {
	TxHash string `form:"txHash" binding:"required"`
}

// TransactionResp .
type TransactionResp struct {
	TxHash string               `json:"tx_hash"`
	From   string               `json:"from"`
	To     string               `json:"to"`
	Nonce  uint64               `json:"nonce"`
	Data   string               `json:"data"`
	Value  string               `json:"value"`
	Logs   []TransactionLogResp `json:"logs"`
}

// TransactionLogResp .
type TransactionLogResp struct {
	Index uint   `json:"index"`
	Data  string `json:"data"`
}

// GetTransaction .
func (s *Service) GetTransaction(ctx context.Context, params *GetTransactionReq) (TransactionResp, error) {
	tx, err := s.dao.GetTransactionByHash(ctx, params.TxHash)
	if err != nil {
		return TransactionResp{}, err
	}
	txResp := TransactionResp{
		TxHash: tx.HashHex,
		From:   tx.FromHex,
		To:     tx.ToHex,
		Nonce:  tx.Nonce,
		Data:   string(tx.TxData),
		Value:  tx.Value,
	}
	if len(tx.Logs) > 0 {
		logs := make([]TransactionLogResp, len(tx.Logs))
		for i := range tx.Logs {
			logs[i] = TransactionLogResp{
				Index: tx.Logs[i].LogIndex,
				Data:  string(tx.Logs[i].LogData),
			}
		}
		txResp.Logs = logs
	}
	return txResp, nil
}
