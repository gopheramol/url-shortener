package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"url-shortener/internal/app/shortner/delivery/web"
	"url-shortener/internal/app/shortner/repository"
	"url-shortener/internal/app/shortner/usecase/usecaseimpl"
)

type Server struct {
	httpServer *http.Server
}

func NewServer() *Server {
	serverMux := http.NewServeMux()
	repo := repository.NewInMemoryDB()
	usecase := usecaseimpl.NewShortnerUseCase(repo)
	web.RegisterHandlers(serverMux, usecase)

	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: serverMux,
	}

	return &Server{
		httpServer: httpServer,
	}
}

func (s *Server) Start() {
	go func() {
		log.Printf("Starting server at %s\n", s.httpServer.Addr)
		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not listen on %s: %v\n", s.httpServer.Addr, err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Server is shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.httpServer.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v\n", err)
	}

	log.Println("Server stopped")
}

func (s *Server) Stop() {
	if err := s.httpServer.Close(); err != nil {
		log.Fatalf("Could not stop server gracefully: %v\n", err)
	}
}
