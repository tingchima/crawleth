// Package ethereum provides
package ethereum

import (
	"context"
	"fmt"
	"math/big"

	"portto_interview/indexer-service/internal/crawler"
	"portto_interview/indexer-service/internal/model"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"
)

// EthCrawler .
type EthCrawler struct {
	signer    types.EIP155Signer
	ethclient *ethclient.Client
}

// NewEthCrawler .
func NewEthCrawler(signer types.EIP155Signer, ethclient *ethclient.Client) crawler.Crawler {
	return &EthCrawler{
		signer:    signer,
		ethclient: ethclient,
	}
}

// HeaderNumber .
func (e *EthCrawler) HeaderNumber(ctx context.Context) (*big.Int, error) {
	header, err := e.ethclient.HeaderByNumber(ctx, nil) // input nil val will return the lastest block's header
	if err != nil {
		return nil, err
	}
	return header.Number, nil
}

// CurrentBlockNumber .
func (e *EthCrawler) CurrentBlockNumber(ctx context.Context) (*big.Int, error) {
	lastBlockNumber, err := e.ethclient.BlockNumber(ctx)
	if err != nil {
		return big.NewInt(0), err
	}
	return big.NewInt(int64(lastBlockNumber)), nil
}

// BlockByNumber .
func (e *EthCrawler) BlockByNumber(ctx context.Context, number *big.Int) (model.EthBlock, error) {
	blockResp, err := e.ethclient.BlockByNumber(ctx, number)
	if err != nil {
		return model.EthBlock{}, err
	}
	return model.EthBlock{
		Number:        decimal.NewFromBigInt(blockResp.Number(), 0),
		HashHex:       blockResp.Hash().Hex(),
		Time:          int64(blockResp.Time()),
		ParentHashHex: blockResp.ParentHash().Hex(),
	}, nil
}

// TransactionsByBlockHashHex .
func (e *EthCrawler) TransactionsByBlockHashHex(ctx context.Context, blockHashHex string) ([]model.EthTransaction, error) {
	txCnt, err := e.ethclient.TransactionCount(ctx, common.HexToHash(blockHashHex))
	if err != nil {
		return nil, err
	}
	transactions := make([]model.EthTransaction, txCnt)

	for txIndex := 0; txIndex < int(txCnt); txIndex++ {
		txResp, err := e.ethclient.TransactionInBlock(ctx, common.HexToHash(blockHashHex), uint(txIndex))
		if err != nil {
			fmt.Printf("TransactionInBlock err: %v\n", err)
			continue
		}
		if txResp == nil {
			continue
		}

		message, err := txResp.AsMessage(e.signer, nil)
		if err != nil {
			fmt.Printf("AsMessage err: %v\n", err)
			continue
		}

		tx := model.EthTransaction{
			BlockHashHex: blockHashHex,
			HashHex:      txResp.Hash().Hex(),
			FromHex:      message.From().Hex(),
			Nonce:        message.Nonce(),
			TxData:       message.Data(),
		}
		if message.To() != nil {
			tx.ToHex = message.To().Hex()
		}
		if message.Value() != nil {
			tx.Value = message.Value().String()
		}

		txLogs, err := e.TransactionLogsByTxHashHex(ctx, txResp.Hash().Hex())
		if err != nil {
			fmt.Printf("TransactionLogsByTxHashHex err: %v\n", err)
			continue
		}
		tx.Logs = txLogs

		transactions[txIndex] = tx
	}

	return transactions, nil
}

// TransactionLogsByTxHashHex .
func (e *EthCrawler) TransactionLogsByTxHashHex(ctx context.Context, txHashHex string) ([]model.EthTransactionLog, error) {
	receiptResp, err := e.ethclient.TransactionReceipt(ctx, common.HexToHash(txHashHex))
	if err != nil {
		return nil, err
	}
	txLogs := make([]model.EthTransactionLog, len(receiptResp.Logs))
	for i, log := range receiptResp.Logs {
		txLogs[i] = model.EthTransactionLog{
			LogIndex:  log.Index,
			LogData:   log.Data,
			TxHashHex: txHashHex,
		}
	}
	return txLogs, nil
}
