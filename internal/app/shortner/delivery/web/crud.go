package web

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"url-shortener/internal/app/shortner/usecase"
)

type ShortnerHandler struct {
	usecase usecase.ShortnerUseCase
}

func NewShortnerHandler(usecase usecase.ShortnerUseCase) *ShortnerHandler {
	return &ShortnerHandler{
		usecase: usecase,
	}
}

func (h *ShortnerHandler) CreateShortURL(w http.ResponseWriter, r *http.Request) {
	log.Println("ShortnerHandler CreateShortURL()")
	// Parse the request body to get the full URL
	var url usecase.ShortURL
	err := json.NewDecoder(r.Body).Decode(&url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Call the EncodeURL method of the use case
	urlData, err := h.usecase.EncodeURL(r.Context(), url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write the short URL ID as response
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s", urlData.ShortURL)
}

func (h *ShortnerHandler) GetOriginalURL(w http.ResponseWriter, r *http.Request) {
	log.Println("ShortnerHandler GetOriginalURL()")
	// Extract the encoded URL parameter from the request URL path
	urlParts := strings.Split(r.URL.Path, "/")
	encodedURL := urlParts[len(urlParts)-1]

	// Decode the encoded URL using your use case layer or repository
	ctx := context.Background()
	originalURL, err := h.usecase.DecodeURL(ctx, encodedURL)
	if err != nil {
		http.Error(w, "URL decoding failed", http.StatusInternalServerError)
		return
	}

	// Return the original URL as a JSON response
	response := struct {
		OriginalURL string `json:"originalURL"`
	}{
		OriginalURL: originalURL.URL,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to serialize response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)

}
