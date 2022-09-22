// Package handler provides
package handler

import (
	"context"
	"fmt"
	"time"

	"portto_interview/indexer-service/global"
	"portto_interview/indexer-service/internal/service"
	"portto_interview/indexer-service/internal/subscriber"
)

// Handler .
type Handler struct {
	service *service.Service
}

// NewHandler .
func NewHandler() *Handler {
	return &Handler{service: service.NewService()}
}

// RunEthereumHeaderSyncing 同步鏈上最新區塊
func (h *Handler) RunEthereumHeaderSyncing(ctx context.Context) {
	// 建立訂閱事件: 這邊只是模擬鏈上websocket服務，一直收最新區塊
	ethSub := subscriber.NewEthHeader(global.EthereumConfig.SubscriberDuration)
	ethSub.Subscribe(ctx, h.service.SubscribeEthHeader)

	for {
		select {
		case <-ctx.Done():

		case currentBlockNumber := <-ethSub.Headers():
			if currentBlockNumber == nil {
				continue
			}
			fmt.Printf("start to sync ethereum header: %v...\n", currentBlockNumber)
			err := h.service.SyncEthereumHeader(ctx, currentBlockNumber)
			if err != nil {
				fmt.Printf("SyncEthereumHeader err: %v\n", err)
				continue
			}
		}
	}
}

// RunEthereumSyncing .
func (h *Handler) RunEthereumSyncing(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():

		default:
			time.Sleep(time.Second * global.AppConfig.EthSyncingDuration)
			fmt.Println("start to sync ethereum block and transactions...")

			err := h.service.SyncEthereum(ctx)
			if err != nil {
				fmt.Printf("SyncEthereum err: %v\n", err)
				continue
			}
		}
	}
}

// RunEthereumBlockStable .
func (h *Handler) RunEthereumBlockStable(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():

		default:
			time.Sleep(time.Second * global.AppConfig.EthUpdateBlockStatus)
			fmt.Println("start to update ethereum block status...")

			err := h.service.UpdateEthBlockStable(ctx)
			if err != nil {
				fmt.Printf("UpdateEthBlockStable err: %v\n", err)
				continue
			}
		}
	}
}
