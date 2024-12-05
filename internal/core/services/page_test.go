package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChooseRandomPage(t *testing.T) {
	service := NewPageService(10)
	sentPages := []int{1, 2, 3}

	page, err := service.ChooseRandom(service.pages, sentPages)

	assert.NoError(t, err)
	assert.NotContains(t, sentPages, page)
	assert.Contains(t, service.pages, page)
}

func TestChooseRandomPage_NoAvailablePages(t *testing.T) {
	service := NewPageService(0)
	page, err := service.ChooseRandom(service.pages, []int{})

	assert.Error(t, err)
	assert.Equal(t, 0, page)
}
