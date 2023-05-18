package usecase

import (
	"context"
)

type ShortnerUseCase interface {
	EncodeURL(ctx context.Context, url ShortURL) (ShortURL, error)
	DecodeURL(ctx context.Context, url string) (ShortURL, error)
}
