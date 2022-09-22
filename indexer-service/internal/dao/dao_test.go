// Package dao provides
package dao

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"testing"

	"portto_interview/indexer-service/global"
	"portto_interview/indexer-service/internal/infra"
	"portto_interview/indexer-service/internal/model"

	"github.com/shopspring/decimal"
)

func init() {
	err := infra.SetupInfra()
	if err != nil {
		log.Fatal(err)
	}
}

// TestAddBlockNumbers .
func TestAddBlockNumbers(t *testing.T) {
	ctx := context.Background()
	dao := NewDao(global.Database, global.Cache)
	err := dao.AddBlockNumbers(ctx, "eth_block_numbers", big.NewInt(9999999999))
	if err != nil {
		t.Error(err)
	}
}

// TestGetBlockNumbers .
func TestGetBlockNumbers(t *testing.T) {
	ctx := context.Background()
	dao := NewDao(global.Database, global.Cache)
	numbers, err := dao.GetBlockNumbers(ctx, "eth_block_number")
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("numbers: %v\n", numbers)
}

// Test .
func TestBatchCrate(t *testing.T) {
	ctx := context.Background()
	dao := NewDao(global.Database, global.Cache)

	blocks := []model.EthBlock{
		{
			Number:  decimal.NewFromBigInt(big.NewInt(8888888), 0),
			HashHex: "123",
		},
		{
			Number:  decimal.NewFromBigInt(big.NewInt(7777777), 0),
			HashHex: "456",
		},
	}

	err := dao.BatchCreateEthBlock(ctx, blocks)
	if err != nil {
		t.Error(err)
	}
}
