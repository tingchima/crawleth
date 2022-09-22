// Package service provides
package service

import (
	"context"

	"portto_interview/api-service/pkg/convert"

	"github.com/shopspring/decimal"
)

// GetBlockListReq .
type GetBlockListReq struct {
	Limit uint `form:"limit" binding:"required"`
}

// GetBlockReq .
type GetBlockReq struct {
	ID uint64 `form:"id" binding:"required"`
}

// BlockResp .
type BlockResp struct {
	BlockNum     decimal.Decimal `json:"block_num"`
	BlockHash    string          `json:"block_hash"`
	BlockTime    int64           `json:"block_time"`
	ParentHash   string          `json:"parent_hash"`
	IsStable     bool            `json:"is_stable"`
	Transactions []string        `json:"transactions,omitempty"`
}

// GetBlockList .
func (s *Service) GetBlockList(ctx context.Context, params *GetBlockListReq) ([]BlockResp, error) {
	blocks, err := s.dao.GetBlockListByLimit(ctx, params.Limit)
	if err != nil {
		return nil, err
	}
	blockResp := make([]BlockResp, len(blocks))
	for i := range blocks {
		blockResp[i] = BlockResp{
			BlockNum:   blocks[i].Number,
			BlockHash:  blocks[i].HashHex,
			BlockTime:  blocks[i].Time,
			ParentHash: blocks[i].ParentHashHex,
			IsStable:   convert.IntTo(blocks[i].IsStable).Bool(),
		}
	}
	return blockResp, nil
}

// GetBlock .
func (s *Service) GetBlock(ctx context.Context, params *GetBlockReq) (BlockResp, error) {
	block, err := s.dao.GetBlockByNumber(ctx, params.ID)
	if err != nil {
		return BlockResp{}, err
	}
	blockResp := BlockResp{
		BlockNum:   block.Number,
		BlockHash:  block.HashHex,
		BlockTime:  block.Time,
		ParentHash: block.ParentHashHex,
		IsStable:   convert.IntTo(block.IsStable).Bool(),
	}
	if len(block.Transactions) > 0 {
		txs := make([]string, len(block.Transactions))
		for i := range block.Transactions {
			txs[i] = block.Transactions[i].HashHex
		}
		blockResp.Transactions = txs
	}
	return blockResp, nil
}
