package usecase

import (
	"context"
)

type ShortnerUseCase interface {
	CreateShortURL(ctx context.Context, url ShortURL) (int, error)
	GetLongURL(ctx context.Context, id int) (ShortURL, error)
}
