package usecases

import (
	"context"
	"log"
	"strconv"

	"bookNotification/internal/core/services"
	"bookNotification/internal/infra"
)

type NotifyUseCase struct {
	PageService services.PageService
	S3Storage     *infra.S3Storage
	SNSNotifier   *infra.SNSService
}

func NewNotifyUseCase(pageService services.PageService, s3Storage *infra.S3Storage, snsNotifier *infra.SNSService) *NotifyUseCase {
	return &NotifyUseCase{
		PageService: pageService,
		S3Storage:     s3Storage,
		SNSNotifier:   snsNotifier,
	}
}

func (uc *NotifyUseCase) Execute(ctx context.Context) {
	sentPages := uc.S3Storage.LoadPages(ctx)

	totalPages := 143

	availablePages := make([]int, 0, totalPages)
	for i := 1; i <= totalPages; i++ {
		availablePages = append(availablePages, i)
	}

	page, err := uc.PageService.ChooseRandom(availablePages, sentPages)
	if err != nil {
		log.Fatalf("Fail to choose page: %v", err)
	}

	if page == -1 {
		log.Println("All pages have been sent.")
		return
	}

	message := "A página do dia é: " + strconv.Itoa(page)
	uc.SNSNotifier.PublishMessage(ctx, message)

	sentPages = append(sentPages, page)
	uc.S3Storage.SavePages(ctx, sentPages)
}
