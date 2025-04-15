package main

import (
	"github.com/rocky114/craftman/internal/app"
	"github.com/rocky114/craftman/internal/app/config"
	"log"
)

func main() {
	cfg, err := config.Load("config.yml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	application, err := app.NewApplication(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize application: %v", err)
	}

	application.Start()
}
