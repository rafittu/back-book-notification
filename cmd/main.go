package main

import (
	"context"
	"log"
	"time"

	"bookNotification/config"
	"bookNotification/internal/core/services"
	"bookNotification/internal/infra"
	"bookNotification/internal/usecases"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	s3Storage, err := infra.NewS3Storage(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize S3 storage: %v", err)
	}

	snsNotifier, err := infra.NewSNSService(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize SNS service: %v", err)
	}

	pageService := services.NewPageService(cfg.TotalPages)

	notifyUseCase, err := usecases.NewNotifyUseCase(*pageService, s3Storage, snsNotifier, cfg)
	if err != nil {
		log.Fatalf("Failed to create NotifyUseCase: %v", err)
	}

	notifyUseCase.Execute(ctx)
}
