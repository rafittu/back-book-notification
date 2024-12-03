package service

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

func (s *PageService) ChooseRandom() (int, error) {
	if len(s.pages) == 0 {
		return 0, errors.New("page list is empty")
	}

	randSource := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomIndex := randSource.Intn(len(s.pages))

	randomNumber := s.pages[randomIndex]
	s.pages = append(s.pages[:randomIndex], s.pages[randomIndex+1:]...)

	return randomNumber, nil
}
