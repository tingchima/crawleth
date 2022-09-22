// Package crawler provides
package crawler

import (
	"context"
	"math/big"
	"sort"
	"sync"

	"portto_interview/indexer-service/global"
	"portto_interview/indexer-service/internal/model"

	"github.com/panjf2000/ants/v2"
)

// Crawler .
type Crawler interface {
	HeaderNumber(ctx context.Context) (*big.Int, error)

	BlockByNumber(ctx context.Context, number *big.Int) (block model.EthBlock, err error)

	TransactionsByBlockHashHex(ctx context.Context, hashHex string) ([]model.EthTransaction, error)

	TransactionLogsByTxHashHex(ctx context.Context, txHashHex string) ([]model.EthTransactionLog, error)

	CurrentBlockNumber(ctx context.Context) (*big.Int, error)
}

type crawlTaskFunc func(ctx context.Context, blockNumber *big.Int) (model.EthBlock, []model.EthTransaction, []model.EthTransactionLog, error)

// TaskWorker .
type TaskWorker struct {
	pool  *ants.Pool
	wg    sync.WaitGroup
	tasks []*big.Int
	cache model.EthCache
	err   error
}

// NewCrawlerTask .
func NewCrawlerTask(tasks []*big.Int) *TaskWorker {
	return &TaskWorker{
		pool:  global.EthCrawlerPool,
		tasks: tasks,
	}
}

// HasErr .
func (tp *TaskWorker) HasErr() bool {
	return tp.err != nil
}

// Err .
func (tp *TaskWorker) Err() error {
	return tp.err
}

// Dispatch .
func (tp *TaskWorker) Dispatch(ctx context.Context, f crawlTaskFunc) {
	for i := range tp.tasks {
		tp.wg.Add(1)
		tp.pool.Submit(tp.crawlFuncWrapper(ctx, &tp.wg, tp.tasks[i], f))
	}
	tp.wg.Wait()
}

// LastBlockNumber .
func (tp *TaskWorker) LastBlockNumber() (*big.Int, bool) {
	if len(tp.tasks) <= 0 {
		return nil, false
	}
	sort.Slice(tp.tasks, func(i, j int) bool {
		a := tp.tasks[i]
		b := tp.tasks[j]
		return a.Int64() > b.Int64()
	})
	return tp.tasks[0], true
}

// Blocks .
func (tp *TaskWorker) Blocks() []model.EthBlock {
	return tp.cache.GetBlocks()
}

// Transactions .
func (tp *TaskWorker) Transactions() []model.EthTransaction {
	return tp.cache.GetTransactions()
}

// TransactionLogs .
func (tp *TaskWorker) TransactionLogs() []model.EthTransactionLog {
	return tp.cache.GetTransactionLogs()
}

// crawlFuncWrapper .
func (tp *TaskWorker) crawlFuncWrapper(ctx context.Context, wg *sync.WaitGroup, task *big.Int, f crawlTaskFunc,
) func() {
	return func() {
		defer wg.Done()
		block, txs, txLogs, err := f(ctx, task)
		if err != nil {
			tp.err = err
			return
		}
		tp.cache.SetBlock(block)
		tp.cache.SetTransactions(txs)
		tp.cache.SetTxLogs(txLogs)
		return
	}
}
