package repository

import (
	"context"
	"errors"
	"log"
	"sync"
	"url-shortener/internal/app/shortner/usecase"
)

type ShortnerRepository interface {
	CreateShortURL(ctx context.Context, url usecase.ShortURL) (string, error)
	GetShortURL(ctx context.Context, id string) (usecase.ShortURL, error)
}

type inMemoryDB struct {
	sync.RWMutex
	urls map[int]usecase.ShortURL
}

func NewInMemoryDB() *inMemoryDB {
	return &inMemoryDB{
		urls: make(map[int]usecase.ShortURL),
	}
}

func (db *inMemoryDB) CreateShortURL(ctx context.Context, url usecase.ShortURL) (int, error) {
	log.Println("inMemoryDB CreateShortURL()")
	log.Println(url.ID, url.FullURL)
	db.Lock()
	defer db.Unlock()

	if _, ok := db.urls[url.ID]; ok {
		return 0, errors.New("short URL with given ID already exists")
	}

	db.urls[url.ID] = url
	return url.ID, nil
}

func (db *inMemoryDB) GetLongURL(ctx context.Context, id int) (usecase.ShortURL, error) {
	log.Println("inMemoryDB GetLongURL()")

	db.RLock()
	defer db.RUnlock()

	if url, ok := db.urls[id]; ok {
		return url, nil
	}

	return usecase.ShortURL{}, errors.New("short URL not found")
}
