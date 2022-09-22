// Package main provides
package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"portto_interview/indexer-service/internal/handler"
	"portto_interview/indexer-service/internal/infra"
	"runtime"
	"syscall"
)

// init .
func init() {
	err := infra.SetupInfra()
	if err != nil {
		log.Fatalf("setup infra err: %v", err)
	}
}

func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

var incomingURLs = []string{"test1", "test2", "test3"}

func main() {
	h := handler.NewHandler()

	headerSyncingCtx, headerSyncingCancel := context.WithCancel(context.Background())
	go func() {
		defer processRecover()
		h.RunEthereumHeaderSyncing(headerSyncingCtx)
	}()

	ethereumSyncingCtx, ethereumSyncingCancel := context.WithCancel(context.Background())
	go func() {
		defer processRecover()
		h.RunEthereumSyncing(ethereumSyncingCtx)
	}()

	ethereumBlockStableCtx, ethereumBlockStableCancel := context.WithCancel(context.Background())
	go func() {
		defer processRecover()
		h.RunEthereumBlockStable(ethereumBlockStableCtx)
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	headerSyncingCancel()
	ethereumSyncingCancel()
	ethereumBlockStableCancel()
	fmt.Println("shutting down server ...")
}

func processRecover() {
	if r := recover(); r != nil {
		var msg string
		for i := 2; ; i++ {
			_, file, line, ok := runtime.Caller(i)
			if !ok {
				break
			}
			msg = msg + fmt.Sprintf("%s:%d\n", file, line)
		}
		fmt.Printf("%s\n↧↧↧↧↧↧ PANIC ↧↧↧↧↧↧\n%s↥↥↥↥↥↥ PANIC ↥↥↥↥↥↥", r, msg)
	}
}
