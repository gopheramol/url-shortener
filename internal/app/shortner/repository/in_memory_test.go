package repository

import (
	"context"
	"testing"
	"url-shortener/internal/app/shortner/usecase"
)

func TestSave(t *testing.T) {
	// Create a new instance of inMemoryDB
	db := NewInMemoryDB().(*inMemoryDB)

	// Create a context
	ctx := context.Background()

	// Define a sample ShortURL
	shortURL := usecase.ShortURL{
		ShortURL: "abc123",
		URL:      "http://google.com",
	}

	// Call the Save method
	result, err := db.Save(ctx, shortURL)

	// Check for errors
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Check the returned result
	if result != shortURL {
		t.Errorf("Unexpected result. Expected: %+v, Got: %+v", shortURL, result)
	}

	// Check if the URL is saved in the map
	db.RLock()
	defer db.RUnlock()
	if db.urls[result.ShortURL] != result {
		t.Errorf("URL not saved in the map")
	}
}

func TestGet(t *testing.T) {
	// Create a new instance of inMemoryDB
	db := NewInMemoryDB().(*inMemoryDB)

	// Create a context
	ctx := context.Background()

	// Define a sample ShortURL
	shortURL := usecase.ShortURL{
		ShortURL: "abc123",
		URL:      "http://google.com",
	}

	// Save the ShortURL in the map
	db.Lock()
	db.urls[shortURL.ShortURL] = shortURL
	db.Unlock()

	// Call the Get method
	result, err := db.Get(ctx, shortURL.ShortURL)

	// Check for errors
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Check the returned result
	if result != shortURL {
		t.Errorf("Unexpected result. Expected: %+v, Got: %+v", shortURL, result)
	}
}
