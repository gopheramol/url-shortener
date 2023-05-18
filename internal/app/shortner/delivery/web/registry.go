package web

import (
	"net/http"
	"url-shortener/internal/app/shortner/usecase"
)

func RegisterHandlers(router *http.ServeMux, usecase usecase.ShortnerUseCase) {
	shortenerHandler := NewShortnerHandler(usecase)

	// Register the handler for creating short URLs
	router.HandleFunc("/encode", shortenerHandler.CreateShortURL)

	// Register the handler for decoding short URLs
	router.HandleFunc("/decode/", shortenerHandler.GetOriginalURL)
}
