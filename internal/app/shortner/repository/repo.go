package repository

import (
	"context"
	"url-shortener/internal/app/shortner/usecase"
)

type ShortnerRepository interface {
	Save(ctx context.Context, url usecase.ShortURL) (int, error)
	Get(ctx context.Context, id int) (usecase.ShortURL, error)
}
