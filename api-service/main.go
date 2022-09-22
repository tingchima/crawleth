package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"portto_interview/api-service/global"
	"portto_interview/api-service/internal/infra"
	"portto_interview/api-service/internal/router"
	"portto_interview/api-service/pkg/config"
)

func init() {
	err := infra.SetupInfra()
	if err != nil {
		log.Fatalf("setup infra err: %v\n", err)
	}
}

//@title API-Service
//@version 1.0
func main() {
	router := router.NewRouter() // self-define routers
	server := &http.Server{
		Addr:           fmt.Sprintf(":%s", global.HTTPServerConfig.Port),
		Handler:        router,
		MaxHeaderBytes: 1 << 20,
	}
	go func() {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal("Listen and serve err")
		}
	}()
	fmt.Printf("http server listen on %s port", global.HTTPServerConfig.Port)

	// Graceful shoutdown
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	err := server.Shutdown(ctx)
	if err != nil {
		log.Fatalf("server forced to shutdown err: %v", err)
	}
	log.Println("shutting down server ...")
}

// setupConfig .
func setupConfig() error {
	cfg, err := config.NewConfig("app", "configs/")
	if err != nil {
		return err
	}
	err = cfg.ReadConfig("App", &global.AppConfig)
	if err != nil {
		return err
	}
	err = cfg.ReadConfig("HTTPServer", &global.HTTPServerConfig)
	if err != nil {
		return err
	}
	// Add new config at here...
	return nil
}
