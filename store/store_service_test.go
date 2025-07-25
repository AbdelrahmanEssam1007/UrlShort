package store

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testStoreService = &StoreService{}

func init() {
	testStoreService = StoreInit()
}

func TestStoreInit(t *testing.T) {
	assert.True(t, testStoreService != nil, "StoreInit should not return nil")
}

func TestInsertionAndRetrieval(t *testing.T) {
	initialLink := "https://www.guru3d.com/news-story/spotted-ryzen-threadripper-pro-3995wx-processor-with-8-channel-ddr4,2.html"
	userUUId := "e0dba740-fc4b-4977-872c-d360239e6b1a"
	shortURL := "Jsz4k57oAX"

	// Persist data mapping
	SaveUrlMapping(shortURL, initialLink, userUUId)

	// Retrieve initial URL
	retrievedUrl, _ := RetrieveInitialUrl(shortURL)

	assert.Equal(t, initialLink, retrievedUrl)
}
