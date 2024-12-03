package main

import (
	"bookNotification/internal/core/services"
	"bookNotification/internal/infra"
	"bookNotification/internal/usecases"
	"context"
)

func main() {
	s3Storage := infra.NewS3Storage()
	snsNotifier := infra.NewSNSService()

	pageService := services.NewPageService()

	notifyUseCase := usecases.NewNotifyUseCase(*pageService, s3Storage, snsNotifier)

	notifyUseCase.Execute(context.TODO())
}
