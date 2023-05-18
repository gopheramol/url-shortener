package web

import (
	"context"
	"fmt"
	"log"
	"net/http"
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
	h.usecase.EncodeURL(context.Background(), usecase.ShortURL{})
}

func (h *ShortnerHandler) GetLongURL(w http.ResponseWriter, r *http.Request) {
	log.Println("ShortnerHandler GetLongURL()")
	data, err := h.usecase.DecodeURL(context.Background(), 1)
	if err != nil {
		log.Println("error while create short url", err.Error())
	}
	fmt.Println(data)
	// Implement get long URL handler logic
}
