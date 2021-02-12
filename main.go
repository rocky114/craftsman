package main

import (
	"context"
	"craftsman/cache"
	"craftsman/config"
	"craftsman/model"
	"craftsman/router"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func init() {
	config.Bootstrap()
	cache.Bootstrap()
	model.Bootstrap()
	router.Bootstrap()
}

func main() {
	fmt.Println("application starting...")

	endPoint := config.GlobalConfig.Server.Addr + ":" + config.GlobalConfig.Server.Port

	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        router.Router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: maxHeaderBytes,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("http server listening %s", endPoint)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	fmt.Println("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		fmt.Println("server forced to shutdown:", err)
	}

	fmt.Println("application exiting...")
}
