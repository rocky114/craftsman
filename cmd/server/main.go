package main

import (
	"github.com/rocky114/craftman/internal/app"
	"github.com/rocky114/craftman/internal/app/config"
	"log"
	"os"
)

func main() {
	configPath := os.Getenv("CRAFTSMAN_CONFIG_PATH")
	if configPath == "" {
		configPath = "config.yaml"
	}
	cfg, err := config.Load(configPath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	application, err := app.NewApplication(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize application: %v", err)
	}

	application.Start()
}
