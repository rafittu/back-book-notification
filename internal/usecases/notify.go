package usecases

import (
	"context"
	"fmt"
	"log"

	"bookNotification/config"
	"bookNotification/internal/core/services"
	"bookNotification/internal/infra"
)

type NotifyUseCase struct {
	PageService services.PageService
	S3Storage     *infra.S3Storage
	SNSNotifier   *infra.SNSService
	Config      *config.Config
}

func NewNotifyUseCase(pageService services.PageService, s3Storage *infra.S3Storage, snsNotifier *infra.SNSService, cfg *config.Config) (*NotifyUseCase, error) {
	if s3Storage == nil || snsNotifier == nil || cfg == nil {
		return nil, fmt.Errorf("dependencies cannot be nil")
	}

	return &NotifyUseCase{
		PageService: pageService,
		S3Storage:   s3Storage,
		SNSNotifier: snsNotifier,
		Config:      cfg,
	}, nil
}

func (uc *NotifyUseCase) Execute(ctx context.Context) {
	sentPages := uc.S3Storage.LoadPages(ctx)

	availablePages := make([]int, 0, uc.Config.TotalPages)
	for i := 1; i <= uc.Config.TotalPages; i++ {
		availablePages = append(availablePages, i)
	}

	page, err := uc.PageService.ChooseRandom(availablePages, sentPages)
	if err != nil {
		log.Fatalf("Fail to choose page: %v", err)
	}

	message := fmt.Sprintf("A página do dia é: %d", page)
	if err := uc.SNSNotifier.PublishMessage(ctx, message); err != nil {
		log.Fatalf("Fail to publish message: %v", err)
	}

	sentPages = append(sentPages, page)
	uc.S3Storage.SavePages(ctx, sentPages)
}
