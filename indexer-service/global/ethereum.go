// Package global provides
package global

import (
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/panjf2000/ants/v2"
)

var (
	// EthereumRPC .
	EthereumRPC *ethclient.Client
	// EthCrawlerPool .
	EthCrawlerPool *ants.Pool
	// Ethsigner .
	Ethsigner types.EIP155Signer
)
