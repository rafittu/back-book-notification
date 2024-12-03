package services

import (
	"errors"
	"math/rand"
	"time"
)

type PageService struct {
	pages []int
}

func NewPageService(totalPages int) *PageService {
	return &PageService{
		pages: generateNumbers(totalPages),
	}
}

func generateNumbers(totalPages int) []int {
	var pages []int
	// set number of pages in the book, 143
	for i := 1; i <= totalPages; i++ {
		pages = append(pages, i)
	}
	return pages
}

func toSet(slice []int) map[int]struct{} {
	set := make(map[int]struct{}, len(slice))
	for _, v := range slice {
		set[v] = struct{}{}
	}
	return set
}

func (s *PageService) ChooseRandom(availablePages []int, sentPages []int) (int, error) {
	sentSet := toSet(sentPages)

	remainingPages := make([]int, 0)
	for _, page := range availablePages {
		if _, exists := sentSet[page]; !exists {
			remainingPages = append(remainingPages, page)
		}
	}

	if len(remainingPages) == 0 {
		return 0, errors.New("page list is empty")
	}

	randSource := rand.New(rand.NewSource(time.Now().UnixNano()))

	return remainingPages[randSource.Intn(len(remainingPages))], nil
}
