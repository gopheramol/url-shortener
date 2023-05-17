package web

import (
	"github.com/gorilla/mux"

	"url-shortener/internal/app/shortner/usecase"
)

func RegisterHandlers(router *mux.Router, usecase usecase.ShortnerUseCase) {
	shortenerHandler := NewShortnerHandler(usecase)

	router.HandleFunc("/encode", shortenerHandler.CreateShortURL).Methods("POST")
	router.HandleFunc("/decode/{url}", shortenerHandler.GetLongURL).Methods("GET")
}
