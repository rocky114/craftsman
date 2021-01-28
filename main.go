package main

import (
	"craftsman/config"
	"craftsman/model"
	"craftsman/router"
	"fmt"
	"log"
	"net/http"
	"time"
)

func init() {
	config.Bootstrap()
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

	log.Printf("[info] start http server listening %s", endPoint)

	err := server.ListenAndServe()

	if err != nil {
		fmt.Printf("server err: %s", err)
	}

	fmt.Println("application shutdown...")
}
