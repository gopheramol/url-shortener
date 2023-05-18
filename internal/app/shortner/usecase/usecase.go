package usecase

import (
	"context"
)

type ShortnerUseCase interface {
	EncodeURL(ctx context.Context, url ShortURL) (int, error)
	DecodeURL(ctx context.Context, id int) (ShortURL, error)
}
