package repository

import (
	"context"
	"url-shortener/internal/app/shortner/usecase"
)

type ShortnerRepository interface {
	Save(ctx context.Context, url usecase.ShortURL) (usecase.ShortURL, error)
	Get(ctx context.Context, url string) (usecase.ShortURL, error)
}
