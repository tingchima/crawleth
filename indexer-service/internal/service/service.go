// Package service provides
package service

import (
	"portto_interview/indexer-service/global"
	"portto_interview/indexer-service/internal/crawler"
	"portto_interview/indexer-service/internal/crawler/ethereum"
	"portto_interview/indexer-service/internal/dao"

	"github.com/panjf2000/ants/v2"
)

// Service .
type Service struct {
	dao            *dao.Dao
	ethCrawler     crawler.Crawler
	ethCrawlerPool *ants.Pool
}

// NewService .
func NewService() *Service {
	return &Service{
		dao:            dao.NewDao(global.Database, global.Cache),
		ethCrawler:     ethereum.NewEthCrawler(global.Ethsigner, global.EthereumRPC),
		ethCrawlerPool: global.EthCrawlerPool,
	}
}
