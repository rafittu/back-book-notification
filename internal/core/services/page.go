package services

import (
	"errors"
	"math/rand"
	"time"
)

type PageService struct {
	pages []int
}

func NewPageService() *PageService {
	return &PageService{
		pages: generateNumbers(),
	}
}

func generateNumbers() []int {
	var pages []int
	// set number of pages in the book, 143
	for i := 1; i <= 143; i++ {
		pages = append(pages, i)
	}
	return pages
}

func (s *PageService) ChooseRandom(availablePages []int, sentPages []int) (int, error) {
	remainingPages := []int{}
	for _, page := range availablePages {
		if !contains(sentPages, page) {
			remainingPages = append(remainingPages, page)
		}
	}

	if len(remainingPages) == 0 {
		return 0, errors.New("page list is empty")
	}

	randSource := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomIndex := randSource.Intn(len(remainingPages))

	randomNumber := remainingPages[randomIndex]
	return randomNumber, nil
}

func contains(slice []int, value int) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}
