package web

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"url-shortener/internal/app/shortner/usecase"
	"url-shortener/mock"
)

func TestCreateShortURL(t *testing.T) {
	// Create a mock implementation of the ShortnerUseCase interface
	mockUsecase := &mock.MockShortnerUseCase{}
	// Define the desired behavior of the EncodeURL method
	mockUsecase.EncodeURLFunc = func(ctx context.Context, url usecase.ShortURL) (usecase.ShortURL, error) {
		// Return a mock response
		return usecase.ShortURL{
			ShortURL: "http://short.est/abc123",
		}, nil
	}

	// Create a new instance of the ShortnerHandler with the mock use case
	handler := NewShortnerHandler(mockUsecase)

	// Create a new HTTP request
	req, err := http.NewRequest("POST", "/encode", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to capture the response
	rr := httptest.NewRecorder()

	// Serve the HTTP request using the handler
	handler.CreateShortURL(rr, req)

	// Check the response status code
	if rr.Code != http.StatusCreated {
		t.Errorf("expected status code %d, got %d", http.StatusCreated, rr.Code)
	}

	// Check the response body
	expectedResponse := "http://short.est/abc123"
	if rr.Body.String() != expectedResponse {
		t.Errorf("expected response body '%s', got '%s'", expectedResponse, rr.Body.String())
	}

}
