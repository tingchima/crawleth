// Package service provides
package service

import (
	"context"
	"errors"
	"fmt"
	"math/big"

	"portto_interview/indexer-service/global"
	"portto_interview/indexer-service/internal/crawler"
	"portto_interview/indexer-service/internal/model"
)

const (
	// EthNeedSyncBlockNumbers .
	EthNeedSyncBlockNumbers = "eth_need_sync_block_numbers"
	// EthCurrentBlockNumber .
	EthCurrentBlockNumber = "eth_current_block_number"
	// EthLastCurrentBlockNumber .
	EthLastCurrentBlockNumber = "eth_last_current_block_number"
	// EthLastBlockNumber .
	EthLastBlockNumber = "eth_last_block_number"
)

// SubscribeEthHeader .
func (s *Service) SubscribeEthHeader(ctx context.Context) (*big.Int, error) {
	return s.ethCrawler.HeaderNumber(ctx)
}

// SyncEthereumHeader .
func (s *Service) SyncEthereumHeader(ctx context.Context, currentBlockNumber *big.Int) error {
	lastCurrentBlockNumber, err := s.dao.GetBlockNumber(ctx, EthLastCurrentBlockNumber)
	if err != nil {
		fmt.Printf("GetBlockNumber err %v\n", err)
		return err
	}

	if currentBlockNumber.Cmp(lastCurrentBlockNumber) <= 0 {
		fmt.Println("no block number need to sync")
		return err
	}

	// 計算下次同步流程所需要的區塊號
	start := lastCurrentBlockNumber.Int64()
	end := currentBlockNumber.Int64()
	needSnycSize := end - start
	if needSnycSize > global.AppConfig.MaxSyncHeaderSize {
		end = global.AppConfig.MaxSyncHeaderSize
	}
	blockNumbers := make([]*big.Int, end)
	for i := int64(0); i < end; i++ {
		start++
		blockNumbers[i] = big.NewInt(start)
	}
	err = s.dao.AddBlockNumbers(ctx, EthNeedSyncBlockNumbers, blockNumbers...)
	if err != nil {
		fmt.Printf("AddBlockNumbers err %v\n", err)
		return err
	}

	// 同步鏈上最新區塊號
	err = s.dao.SetBlockNumber(ctx, EthCurrentBlockNumber, currentBlockNumber)
	if err != nil {
		fmt.Printf("SetBlockNumber err %v\n", err)
		return err
	}
	err = s.dao.SetBlockNumber(ctx, EthLastCurrentBlockNumber, big.NewInt(start))
	if err != nil {
		fmt.Printf("SetBlockNumber err %v\n", err)
		return err
	}

	return nil
}

// SyncEthereum .
func (s *Service) SyncEthereum(ctx context.Context) (err error) {
	// Pop block numbers from redis
	blockNumbers, err := s.dao.PopBlockNumbers(ctx, EthNeedSyncBlockNumbers, global.AppConfig.EthCrawlerSize)
	if err != nil {
		fmt.Printf("PopBlockNumbers err %v\n", err)
		return err
	}
	if len(blockNumbers) <= 0 {
		fmt.Println("no block number need to sync")
		return err
	}

	// 初始化任務
	ct := crawler.NewCrawlerTask(blockNumbers)
	// 開始分配任務並同步資料
	ct.Dispatch(ctx, s.crawlEthereumProcess)

	if ct.HasErr() {
		return ct.Err()
	}

	tx := s.dao.Begin(ctx)
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("recover err")
		}
		if err != nil {
			if e := tx.Rollback(); e != nil {
				fmt.Println("rollback fail")
			}
			err = s.dao.AddBlockNumbers(ctx, EthNeedSyncBlockNumbers, blockNumbers...)
			if err != nil {
				fmt.Printf("AddBlockNumbers err: %v\n", err)
			}
			return
		}
		// 提交TX交易
		if e := tx.Commit(); e != nil {
			fmt.Println("commit fail")
		}
	}()

	// 寫入DB
	err = tx.BatchCreateEthBlock(ctx, ct.Blocks())
	if err != nil {
		return err
	}
	err = tx.BatchCreateEthTransaction(ctx, ct.Transactions())
	if err != nil {
		return err
	}
	err = tx.BatchCreateEthTransactionLog(ctx, ct.TransactionLogs())
	if err != nil {
		return err
	}

	// 更新已同步之最後區塊號
	lastBlockNumber, ok := ct.LastBlockNumber()
	if ok {
		originalLastBlockNumber, err := s.dao.GetBlockNumber(ctx, EthLastBlockNumber)
		if err != nil {
			return err
		}
		if originalLastBlockNumber.Int64() > lastBlockNumber.Int64() {
			return nil
		}
		err = s.dao.SetBlockNumber(ctx, EthLastBlockNumber, lastBlockNumber)
		if err != nil {
			return err
		}
	}
	return nil
}

// UpdateEthBlockStable .
func (s *Service) UpdateEthBlockStable(ctx context.Context) error {
	currentBlockNumber, err := s.dao.GetBlockNumber(ctx, EthCurrentBlockNumber)
	if err != nil {
		return err
	}

	err = s.dao.UpdateEthereumBlockStable(ctx, currentBlockNumber)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) crawlEthereumProcess(ctx context.Context, blockNumber *big.Int) (
	b model.EthBlock,
	txs []model.EthTransaction,
	txLogs []model.EthTransactionLog,
	err error,
) {
	currentBlockNumber, err := s.dao.GetBlockNumber(ctx, EthCurrentBlockNumber)
	if err != nil {
		return
	}

	// 爬取區塊資料
	b, err = s.ethCrawler.BlockByNumber(ctx, blockNumber)
	if err != nil {
		fmt.Printf("BlockByNumber err %v\n", err)
		return
	}

	subVal := big.NewInt(0).Sub(currentBlockNumber, b.Number.BigInt())
	if subVal.Cmp(big.NewInt(int64(global.AppConfig.EthStableBlockRule))) > 0 {
		b.IsStable = model.BlockIsStable
	}

	// 爬取區塊交易和log資料
	txs, err = s.ethCrawler.TransactionsByBlockHashHex(ctx, b.HashHex)
	if err != nil {
		fmt.Printf("TransactionsByBlockHashHex err %v\n", err)
		return
	}
	for i := range txs {
		txLogs = append(txLogs, txs[i].Logs...)
	}
	return
}
