package repository

import (
	"context"
	"errors"
	"log"
	"sync"
	"url-shortener/internal/app/shortner/usecase"
)

type inMemoryDB struct {
	sync.RWMutex
	urls map[string]usecase.ShortURL
}

func NewInMemoryDB() ShortnerRepository {
	return &inMemoryDB{
		urls: make(map[string]usecase.ShortURL),
	}
}

func (db *inMemoryDB) Save(ctx context.Context, url usecase.ShortURL) (usecase.ShortURL, error) {
	log.Println("inMemoryDB Save()")
	db.Lock()
	defer db.Unlock()

	db.urls[url.ShortURL] = url
	return url, nil
}

func (db *inMemoryDB) Get(ctx context.Context, url string) (usecase.ShortURL, error) {
	log.Println("inMemoryDB Get()")
	db.RLock()
	defer db.RUnlock()

	shortURL, ok := db.urls[url]
	if !ok {
		return usecase.ShortURL{}, errors.New("short URL not found")
	}
	return shortURL, nil
}
