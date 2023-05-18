package usecaseimpl

import (
	"context"
	"testing"
	"url-shortener/internal/app/shortner/usecase"
	"url-shortener/mock"
)

func TestEncodeURL(t *testing.T) {
	// Create a mock implementation of the ShortnerRepository interface
	mockRepo := &mock.MockShortnerRepository{}
	// Define the desired behavior of the Save method
	mockRepo.SaveFunc = func(ctx context.Context, url usecase.ShortURL) (usecase.ShortURL, error) {
		// Return a mock response
		return usecase.ShortURL{
			ShortURL: "abc123",
		}, nil
	}

	// Create a new instance of the shortnerUseCaseImpl with the mock repository
	usecaseImpl := NewShortnerUseCase(mockRepo)

	// Create a mock context
	ctx := context.Background()

	// Create a mock URL to encode
	url := usecase.ShortURL{
		URL: "http://google.com",
	}

	// Call the EncodeURL method
	resp, err := usecaseImpl.EncodeURL(ctx, url)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Check the response
	expectedShortURL := "http://short.est/qiI5wXYJ"
	if resp.ShortURL != expectedShortURL {
		t.Errorf("unexpected short URL, got: %s, want: %s", resp.ShortURL, expectedShortURL)
	}

}

func TestDecodeURL(t *testing.T) {
	// Create a mock repository
	repo := &mock.MockShortnerRepository{
		GetFunc: func(ctx context.Context, url string) (usecase.ShortURL, error) {
			// Simulate successful retrieval of the short URL
			return usecase.ShortURL{
				ShortURL: url,
			}, nil
		},
	}

	// Create the use case instance with the mock repository
	usecase := NewShortnerUseCase(repo)

	// Create a context
	ctx := context.Background()

	// Define the expected result
	expectedURL := "http://google.com"

	// Call the DecodeURL method
	result, err := usecase.DecodeURL(ctx, "http://short.est/qiI5wXYJ")

	// Check for errors
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Check the decoded URL
	if result.URL != expectedURL {
		t.Errorf("Unexpected decoded URL. Expected: %s, Got: %s", expectedURL, result.URL)
	}
}
